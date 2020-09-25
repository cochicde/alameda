package keycodes

import (
	"github.com/spf13/cobra"
	"prophetstor.com/alameda/datahub/tools/license-utils/pkg/keycodes"
)

var AddKeycodeCmd = &cobra.Command{
	Use:   "add",
	Short: "add keycode",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		keycodes.AddKeycode(keycode)
	},
}

func init() {
	parseAddKeycodeFlag()
}

func parseAddKeycodeFlag() {
	AddKeycodeCmd.Flags().StringVar(&keycode, "keycode", "", "The keycode which will be to added in datahub.")
}
