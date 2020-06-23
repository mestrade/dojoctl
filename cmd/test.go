package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var (
	testCmd = &cobra.Command{
		Use:   "test",
		Short: "manage your tests",
		Long:  `interact with your defectdojo tests`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Echo: " + strings.Join(args, " "))
		},
	}

	testListCmd = &cobra.Command{
		Use:   "list",
		Short: "list tests",
		Args:  cobra.MinimumNArgs(0),
		Run:   testListExec,
	}
)

func testListExec(cmd *cobra.Command, args []string) {

	err := DojoCtx.RetrieveCurrentProductID()
	if err != nil {
		return
	}

	DojoCtx.RetrieveCurrentEngagementID()

	tests, err := DojoCtx.TestList()
	if err != nil {
		panic(err)
	}

	for _, test := range tests {
		test.DisplayShort()
	}
}
