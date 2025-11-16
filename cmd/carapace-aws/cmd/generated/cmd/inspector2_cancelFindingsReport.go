package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector2_cancelFindingsReportCmd = &cobra.Command{
	Use:   "cancel-findings-report",
	Short: "Cancels the given findings report.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector2_cancelFindingsReportCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(inspector2_cancelFindingsReportCmd).Standalone()

		inspector2_cancelFindingsReportCmd.Flags().String("report-id", "", "The ID of the report to be canceled.")
		inspector2_cancelFindingsReportCmd.MarkFlagRequired("report-id")
	})
	inspector2Cmd.AddCommand(inspector2_cancelFindingsReportCmd)
}
