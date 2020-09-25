package app

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"os"
	Keycodes "prophetstor.com/alameda/datahub/tools/license-utils/pkg/keycodes"
	Setup "prophetstor.com/alameda/datahub/tools/license-utils/pkg/setup"
	Utils "prophetstor.com/alameda/datahub/tools/license-utils/pkg/utils"
)

var (
	InteractiveCmd = &cobra.Command{
		Use:   "interactive",
		Short: "interactive interface to management keycode",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			interactive()
		},
	}
)

func interactive() {
	Utils.ClearScreen(false)

	for true {
		prompt := promptui.Select{
			Label: "Select Option",
			Items: []string{"Keycode", "Setup", "Exit"},
		}

		_, result, err := prompt.Run()
		if err != nil {
			fmt.Printf("Invalid input value %v\n", err)
			return
		}

		switch result {
		case "Keycode":
			result, _ := Keycodes.Executor()
			if result == "Back" {
				Utils.ClearScreen(false)
			} else {
				Utils.ClearScreen(true)
			}
		case "Setup":
			Setup.SetDatahubAddress()
			Utils.ClearScreen(false)
		default:
			os.Exit(0)
		}
	}
}
