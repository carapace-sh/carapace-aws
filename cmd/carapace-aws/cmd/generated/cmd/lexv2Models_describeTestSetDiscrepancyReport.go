package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_describeTestSetDiscrepancyReportCmd = &cobra.Command{
	Use:   "describe-test-set-discrepancy-report",
	Short: "Gets metadata information about the test set discrepancy report.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_describeTestSetDiscrepancyReportCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lexv2Models_describeTestSetDiscrepancyReportCmd).Standalone()

		lexv2Models_describeTestSetDiscrepancyReportCmd.Flags().String("test-set-discrepancy-report-id", "", "The unique identifier of the test set discrepancy report.")
		lexv2Models_describeTestSetDiscrepancyReportCmd.MarkFlagRequired("test-set-discrepancy-report-id")
	})
	lexv2ModelsCmd.AddCommand(lexv2Models_describeTestSetDiscrepancyReportCmd)
}
