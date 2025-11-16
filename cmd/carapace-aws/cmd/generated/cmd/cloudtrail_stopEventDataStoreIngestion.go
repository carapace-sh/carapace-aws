package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudtrail_stopEventDataStoreIngestionCmd = &cobra.Command{
	Use:   "stop-event-data-store-ingestion",
	Short: "Stops the ingestion of live events on an event data store specified as either an ARN or the ID portion of the ARN.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudtrail_stopEventDataStoreIngestionCmd).Standalone()

	cloudtrail_stopEventDataStoreIngestionCmd.Flags().String("event-data-store", "", "The ARN (or ID suffix of the ARN) of the event data store for which you want to stop ingestion.")
	cloudtrail_stopEventDataStoreIngestionCmd.MarkFlagRequired("event-data-store")
	cloudtrailCmd.AddCommand(cloudtrail_stopEventDataStoreIngestionCmd)
}
