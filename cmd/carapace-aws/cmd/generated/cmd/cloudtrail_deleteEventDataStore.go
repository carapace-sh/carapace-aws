package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudtrail_deleteEventDataStoreCmd = &cobra.Command{
	Use:   "delete-event-data-store",
	Short: "Disables the event data store specified by `EventDataStore`, which accepts an event data store ARN.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudtrail_deleteEventDataStoreCmd).Standalone()

	cloudtrail_deleteEventDataStoreCmd.Flags().String("event-data-store", "", "The ARN (or the ID suffix of the ARN) of the event data store to delete.")
	cloudtrail_deleteEventDataStoreCmd.MarkFlagRequired("event-data-store")
	cloudtrailCmd.AddCommand(cloudtrail_deleteEventDataStoreCmd)
}
