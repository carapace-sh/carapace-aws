package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudtrail_getEventDataStoreCmd = &cobra.Command{
	Use:   "get-event-data-store",
	Short: "Returns information about an event data store specified as either an ARN or the ID portion of the ARN.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudtrail_getEventDataStoreCmd).Standalone()

	cloudtrail_getEventDataStoreCmd.Flags().String("event-data-store", "", "The ARN (or ID suffix of the ARN) of the event data store about which you want information.")
	cloudtrail_getEventDataStoreCmd.MarkFlagRequired("event-data-store")
	cloudtrailCmd.AddCommand(cloudtrail_getEventDataStoreCmd)
}
