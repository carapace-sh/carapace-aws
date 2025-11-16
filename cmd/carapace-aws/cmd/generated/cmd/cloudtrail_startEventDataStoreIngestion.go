package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudtrail_startEventDataStoreIngestionCmd = &cobra.Command{
	Use:   "start-event-data-store-ingestion",
	Short: "Starts the ingestion of live events on an event data store specified as either an ARN or the ID portion of the ARN.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudtrail_startEventDataStoreIngestionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudtrail_startEventDataStoreIngestionCmd).Standalone()

		cloudtrail_startEventDataStoreIngestionCmd.Flags().String("event-data-store", "", "The ARN (or ID suffix of the ARN) of the event data store for which you want to start ingestion.")
		cloudtrail_startEventDataStoreIngestionCmd.MarkFlagRequired("event-data-store")
	})
	cloudtrailCmd.AddCommand(cloudtrail_startEventDataStoreIngestionCmd)
}
