package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector2_getFindingsReportStatusCmd = &cobra.Command{
	Use:   "get-findings-report-status",
	Short: "Gets the status of a findings report.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector2_getFindingsReportStatusCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(inspector2_getFindingsReportStatusCmd).Standalone()

		inspector2_getFindingsReportStatusCmd.Flags().String("report-id", "", "The ID of the report to retrieve the status of.")
	})
	inspector2Cmd.AddCommand(inspector2_getFindingsReportStatusCmd)
}
