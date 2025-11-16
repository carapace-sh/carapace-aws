package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudtrail_restoreEventDataStoreCmd = &cobra.Command{
	Use:   "restore-event-data-store",
	Short: "Restores a deleted event data store specified by `EventDataStore`, which accepts an event data store ARN.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudtrail_restoreEventDataStoreCmd).Standalone()

	cloudtrail_restoreEventDataStoreCmd.Flags().String("event-data-store", "", "The ARN (or the ID suffix of the ARN) of the event data store that you want to restore.")
	cloudtrail_restoreEventDataStoreCmd.MarkFlagRequired("event-data-store")
	cloudtrailCmd.AddCommand(cloudtrail_restoreEventDataStoreCmd)
}
