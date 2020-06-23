package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var (
	productCmd = &cobra.Command{
		Use:   "product",
		Short: "manage your products",
		Long:  `interact with your defectdojo products`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Echo: " + strings.Join(args, " "))
		},
	}

	productListCmd = &cobra.Command{
		Use:   "list",
		Short: "list products",
		Args:  cobra.MinimumNArgs(0),
		Run:   productListExec,
	}
)

func productExec(cmd *cobra.Command, args []string) {

	fmt.Printf("Print project: %s\n", args[0])

}

func productListExec(cmd *cobra.Command, args []string) {
	products, err := DojoCtx.ProductList()
	if err != nil {
		panic(err)
	}

	for _, product := range products {
		product.DisplayShort()
	}
}
