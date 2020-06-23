package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	uploadCmd = &cobra.Command{
		Use:   "upload report_filename",
		Short: "upload a new report",
		Args:  cobra.MinimumNArgs(1),
		Run:   uploadExec,
	}

	uploadScanType string
)

func uploadExec(cmd *cobra.Command, args []string) {

	DojoCtx.RetrieveCurrentProductID()
	DojoCtx.RetrieveCurrentEngagementID()

	reportName := args[0]
	if len(reportName) == 0 {
		return
	}

	if len(uploadScanType) == 0 {
		return
	}

	err := DojoCtx.Upload(reportName, uploadScanType)
	if err != nil {
		fmt.Printf("Cannot upload report: %v\n", err)
		os.Exit(-1)
	}

}
