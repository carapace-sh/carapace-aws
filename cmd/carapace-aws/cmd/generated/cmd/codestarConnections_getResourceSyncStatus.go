package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codestarConnections_getResourceSyncStatusCmd = &cobra.Command{
	Use:   "get-resource-sync-status",
	Short: "Returns the status of the sync with the Git repository for a specific Amazon Web Services resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codestarConnections_getResourceSyncStatusCmd).Standalone()

	codestarConnections_getResourceSyncStatusCmd.Flags().String("resource-name", "", "The name of the Amazon Web Services resource for the sync status with the Git repository.")
	codestarConnections_getResourceSyncStatusCmd.Flags().String("sync-type", "", "The sync type for the sync status with the Git repository.")
	codestarConnections_getResourceSyncStatusCmd.MarkFlagRequired("resource-name")
	codestarConnections_getResourceSyncStatusCmd.MarkFlagRequired("sync-type")
	codestarConnectionsCmd.AddCommand(codestarConnections_getResourceSyncStatusCmd)
}
