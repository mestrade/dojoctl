package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var (
	engagementCmd = &cobra.Command{
		Use:   "engagement",
		Short: "manage your engagements",
		Long:  `interact with your defectdojo engagements`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Echo: " + strings.Join(args, " "))
		},
	}

	engagementListCmd = &cobra.Command{
		Use:   "list",
		Short: "list items",
		Args:  cobra.MinimumNArgs(0),
		Run:   engagementListExec,
	}
)

func engagementListExec(cmd *cobra.Command, args []string) {

	err := DojoCtx.RetrieveCurrentProductID()
	if err != nil {
		return
	}

	engs, err := DojoCtx.EngagementList()
	if err != nil {
		panic(err)
	}

	for _, eng := range engs {
		eng.DisplayShort()
	}
}
