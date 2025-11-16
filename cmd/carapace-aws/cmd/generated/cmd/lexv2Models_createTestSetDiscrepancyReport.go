package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_createTestSetDiscrepancyReportCmd = &cobra.Command{
	Use:   "create-test-set-discrepancy-report",
	Short: "Create a report that describes the differences between the bot and the test set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_createTestSetDiscrepancyReportCmd).Standalone()

	lexv2Models_createTestSetDiscrepancyReportCmd.Flags().String("target", "", "The target bot for the test set discrepancy report.")
	lexv2Models_createTestSetDiscrepancyReportCmd.Flags().String("test-set-id", "", "The test set Id for the test set discrepancy report.")
	lexv2Models_createTestSetDiscrepancyReportCmd.MarkFlagRequired("target")
	lexv2Models_createTestSetDiscrepancyReportCmd.MarkFlagRequired("test-set-id")
	lexv2ModelsCmd.AddCommand(lexv2Models_createTestSetDiscrepancyReportCmd)
}
