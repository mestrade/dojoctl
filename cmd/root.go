package cmd

import (
	"dojoctl/pkg/dojo"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	DojoCtx *dojo.Ctx

	rootCmd = &cobra.Command{
		Use:   "dojoctl",
		Short: "dojoctl is a cli tool to interact with defectdojo",
		Long: `dojoctl is a cli tool to interact with defectdojo.
You can interact with your products, engagements and tests, or use it to upload scan reports`,
		Args: cobra.MinimumNArgs(1),
		Run:  productExec,
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {

	rootCmd.AddCommand(productCmd)
	productCmd.AddCommand(productListCmd)

	rootCmd.AddCommand(engagementCmd)
	engagementCmd.AddCommand(engagementListCmd)

	rootCmd.AddCommand(testCmd)
	testCmd.AddCommand(testListCmd)

	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(configSetCmd)

	rootCmd.AddCommand(uploadCmd)
	uploadCmd.PersistentFlags().StringVar(&uploadScanType, "type", "", "scan type of the report (ie. ZAP Scan)")

}

func er(msg interface{}) {
	fmt.Println("Error:", msg)
	os.Exit(1)
}
