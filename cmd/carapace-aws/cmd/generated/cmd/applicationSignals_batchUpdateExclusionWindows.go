package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var applicationSignals_batchUpdateExclusionWindowsCmd = &cobra.Command{
	Use:   "batch-update-exclusion-windows",
	Short: "Add or remove time window exclusions for one or more Service Level Objectives (SLOs).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(applicationSignals_batchUpdateExclusionWindowsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(applicationSignals_batchUpdateExclusionWindowsCmd).Standalone()

		applicationSignals_batchUpdateExclusionWindowsCmd.Flags().String("add-exclusion-windows", "", "A list of exclusion windows to add to the specified SLOs.")
		applicationSignals_batchUpdateExclusionWindowsCmd.Flags().String("remove-exclusion-windows", "", "A list of exclusion windows to remove from the specified SLOs.")
		applicationSignals_batchUpdateExclusionWindowsCmd.Flags().String("slo-ids", "", "The list of SLO IDs to add or remove exclusion windows from.")
		applicationSignals_batchUpdateExclusionWindowsCmd.MarkFlagRequired("slo-ids")
	})
	applicationSignalsCmd.AddCommand(applicationSignals_batchUpdateExclusionWindowsCmd)
}
