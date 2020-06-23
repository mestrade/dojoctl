package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	configCmd = &cobra.Command{
		Use:   "config",
		Short: "manage your configuration",
		Long:  `set your defectdojo context`,
		Args:  cobra.MinimumNArgs(0),
		Run:   displayExec,
	}

	configSetCmd = &cobra.Command{
		Use:   "set",
		Short: "set value",
		Args:  cobra.MinimumNArgs(2),
		Run:   configSetExec,
	}
)

func displayExec(cmd *cobra.Command, args []string) {
	fmt.Printf("Use defectdojo API: %s\n", DojoCtx.Setup.ApiBaseUrl)
	fmt.Printf("Current product: %s\n", DojoCtx.Context.CurrentProduct)
	fmt.Printf("Current engagement: %s\n", DojoCtx.Context.CurrentEngagement)
}

func configSetExec(cmd *cobra.Command, args []string) {

	configElem := args[0]
	configValue := args[1]

	if len(configElem) == 0 || len(configValue) == 0 {
		fmt.Printf("Invalid value")
		return
	}

	switch configElem {
	case "product":
		err := DojoCtx.SetProductByName(configValue)
		if err != nil {
			fmt.Printf("Error setting product: %v\n", err)
			return
		}
		DojoCtx.Save()
	case "engagement":

		DojoCtx.RetrieveCurrentProductID()

		err := DojoCtx.SetEngagementByName(configValue)
		if err != nil {
			fmt.Printf("Error setting engagement: %v\n", err)
			return
		}
		DojoCtx.Save()
	}

}
