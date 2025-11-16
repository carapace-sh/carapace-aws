package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeconnections_getSyncBlockerSummaryCmd = &cobra.Command{
	Use:   "get-sync-blocker-summary",
	Short: "Returns a list of the most recent sync blockers.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeconnections_getSyncBlockerSummaryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codeconnections_getSyncBlockerSummaryCmd).Standalone()

		codeconnections_getSyncBlockerSummaryCmd.Flags().String("resource-name", "", "The name of the Amazon Web Services resource currently blocked from automatically being synced from a Git repository.")
		codeconnections_getSyncBlockerSummaryCmd.Flags().String("sync-type", "", "The sync type for the sync blocker summary.")
		codeconnections_getSyncBlockerSummaryCmd.MarkFlagRequired("resource-name")
		codeconnections_getSyncBlockerSummaryCmd.MarkFlagRequired("sync-type")
	})
	codeconnectionsCmd.AddCommand(codeconnections_getSyncBlockerSummaryCmd)
}
