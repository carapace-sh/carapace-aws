package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeconnections_updateSyncBlockerCmd = &cobra.Command{
	Use:   "update-sync-blocker",
	Short: "Allows you to update the status of a sync blocker, resolving the blocker and allowing syncing to continue.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeconnections_updateSyncBlockerCmd).Standalone()

	codeconnections_updateSyncBlockerCmd.Flags().String("id", "", "The ID of the sync blocker to be updated.")
	codeconnections_updateSyncBlockerCmd.Flags().String("resolved-reason", "", "The reason for resolving the sync blocker.")
	codeconnections_updateSyncBlockerCmd.Flags().String("resource-name", "", "The name of the resource for the sync blocker to be updated.")
	codeconnections_updateSyncBlockerCmd.Flags().String("sync-type", "", "The sync type of the sync blocker to be updated.")
	codeconnections_updateSyncBlockerCmd.MarkFlagRequired("id")
	codeconnections_updateSyncBlockerCmd.MarkFlagRequired("resolved-reason")
	codeconnections_updateSyncBlockerCmd.MarkFlagRequired("resource-name")
	codeconnections_updateSyncBlockerCmd.MarkFlagRequired("sync-type")
	codeconnectionsCmd.AddCommand(codeconnections_updateSyncBlockerCmd)
}
