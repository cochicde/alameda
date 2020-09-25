package app

import (
	"errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"prophetstor.com/alameda/cmd/app"
	Keycodes "prophetstor.com/alameda/datahub/pkg/account-mgt/keycodes"
	Licenses "prophetstor.com/alameda/datahub/pkg/account-mgt/licenses"
	DatahubConfig "prophetstor.com/alameda/datahub/pkg/config"
	Notifier "prophetstor.com/alameda/datahub/pkg/notifier"
	SchemaMgt "prophetstor.com/alameda/datahub/pkg/schemamgt"
	EventMgt "prophetstor.com/alameda/internal/pkg/event-mgt"
	"prophetstor.com/alameda/pkg/utils/log"
	"strings"
)

var RootCmd = &cobra.Command{
	Use:   "datahub",
	Short: "alameda datahub",
	Long:  "",
}

var (
	configurationFilePath string

	scope  *log.Scope
	config DatahubConfig.Config
)

func init() {
	RootCmd.AddCommand(RunCmd)
	RootCmd.AddCommand(app.VersionCmd)
	RootCmd.AddCommand(ProbeCmd)

	RootCmd.PersistentFlags().StringVar(&configurationFilePath, "config", "/etc/alameda/datahub/datahub.toml", "The path to datahub configuration file.")
}

func setLoggerScopesWithConfig(config log.Config) {
	for _, scope := range log.Scopes() {
		scope.SetLogCallers(config.SetLogCallers == true)
		if outputLvl, ok := log.StringToLevel(config.OutputLevel); ok {
			scope.SetOutputLevel(outputLvl)
		}
		if stacktraceLevel, ok := log.StringToLevel(config.StackTraceLevel); ok {
			scope.SetStackTraceLevel(stacktraceLevel)
		}
	}
}

func mergeConfigFileValueWithDefaultConfigValue() {
	if configurationFilePath == "" {

	} else {

		viper.SetConfigFile(configurationFilePath)
		err := viper.ReadInConfig()
		if err != nil {
			panic(errors.New("Failed to read configuration file: " + err.Error()))
		}
		err = viper.Unmarshal(&config)
		if err != nil {
			panic(errors.New("Failed to unmarshal configuration file: " + err.Error()))
		}
	}
}

func initConfig() {
	config = DatahubConfig.NewDefaultConfig()
	initViperSetting()
	mergeConfigFileValueWithDefaultConfigValue()
}

func initViperSetting() {
	viper.SetEnvPrefix(envVarPrefix)
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
}

func initLogger() {
	opt := log.DefaultOptions()
	opt.RotationMaxSize = defaultRotationMaxSizeMegabytes
	opt.RotationMaxBackups = defaultRotationMaxBackups
	opt.RotateOutputPath = defaultLogRotateOutputFile
	err := log.Configure(opt)
	if err != nil {
		panic(err)
	}

	scope = log.RegisterScope("datahub_probe", "datahub probe command", 0)
}

func initEventMgt() {
	scope.Info("Initialize event management")
	EventMgt.InitEventMgt(config.InfluxDB, config.RabbitMQ)
}

func initKeycode() {
	scope.Info("Initialize keycode management")
	Keycodes.KeycodeInit(config.Keycode)
	keycodeMgt := Keycodes.NewKeycodeMgt(config.InfluxDB)
	keycodeMgt.Refresh(true)
}

func initLicense() {
	scope.Info("Initialize license management")
	Licenses.LicenseInit(config.License)
}

func initSchema() {
	scope.Info("Initialize schema management")
	SchemaMgt.SchemaInit(config.InfluxDB)
	SchemaMgt.DefaultSchemasInit()
	schemaMgt := SchemaMgt.NewSchemaManagement()
	schemaMgt.Refresh()
}

func initNotifier() {
	scope.Info("Initialize notifier")
	Notifier.Init(config.Notifier, config.InfluxDB)
	go Notifier.Run()
}
