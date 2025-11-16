package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codestarConnections_getSyncBlockerSummaryCmd = &cobra.Command{
	Use:   "get-sync-blocker-summary",
	Short: "Returns a list of the most recent sync blockers.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codestarConnections_getSyncBlockerSummaryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codestarConnections_getSyncBlockerSummaryCmd).Standalone()

		codestarConnections_getSyncBlockerSummaryCmd.Flags().String("resource-name", "", "The name of the Amazon Web Services resource currently blocked from automatically being synced from a Git repository.")
		codestarConnections_getSyncBlockerSummaryCmd.Flags().String("sync-type", "", "The sync type for the sync blocker summary.")
		codestarConnections_getSyncBlockerSummaryCmd.MarkFlagRequired("resource-name")
		codestarConnections_getSyncBlockerSummaryCmd.MarkFlagRequired("sync-type")
	})
	codestarConnectionsCmd.AddCommand(codestarConnections_getSyncBlockerSummaryCmd)
}
