package cmd

import (
	"dojoctl/pkg/dojo"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	DojoCtx *dojo.Ctx
	cfgFile string

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
	cobra.OnInitialize(initConfig)

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

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		if err != nil {
			er(err)
		}

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".cobra")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
