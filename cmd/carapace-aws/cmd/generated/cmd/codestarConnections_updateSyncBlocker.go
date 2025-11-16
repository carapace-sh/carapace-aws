package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codestarConnections_updateSyncBlockerCmd = &cobra.Command{
	Use:   "update-sync-blocker",
	Short: "Allows you to update the status of a sync blocker, resolving the blocker and allowing syncing to continue.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codestarConnections_updateSyncBlockerCmd).Standalone()

	codestarConnections_updateSyncBlockerCmd.Flags().String("id", "", "The ID of the sync blocker to be updated.")
	codestarConnections_updateSyncBlockerCmd.Flags().String("resolved-reason", "", "The reason for resolving the sync blocker.")
	codestarConnections_updateSyncBlockerCmd.Flags().String("resource-name", "", "The name of the resource for the sync blocker to be updated.")
	codestarConnections_updateSyncBlockerCmd.Flags().String("sync-type", "", "The sync type of the sync blocker to be updated.")
	codestarConnections_updateSyncBlockerCmd.MarkFlagRequired("id")
	codestarConnections_updateSyncBlockerCmd.MarkFlagRequired("resolved-reason")
	codestarConnections_updateSyncBlockerCmd.MarkFlagRequired("resource-name")
	codestarConnections_updateSyncBlockerCmd.MarkFlagRequired("sync-type")
	codestarConnectionsCmd.AddCommand(codestarConnections_updateSyncBlockerCmd)
}
