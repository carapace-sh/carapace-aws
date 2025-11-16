package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudtrail_createChannelCmd = &cobra.Command{
	Use:   "create-channel",
	Short: "Creates a channel for CloudTrail to ingest events from a partner or external source.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudtrail_createChannelCmd).Standalone()

	cloudtrail_createChannelCmd.Flags().String("destinations", "", "One or more event data stores to which events arriving through a channel will be logged.")
	cloudtrail_createChannelCmd.Flags().String("name", "", "The name of the channel.")
	cloudtrail_createChannelCmd.Flags().String("source", "", "The name of the partner or external event source.")
	cloudtrail_createChannelCmd.Flags().String("tags", "", "")
	cloudtrail_createChannelCmd.MarkFlagRequired("destinations")
	cloudtrail_createChannelCmd.MarkFlagRequired("name")
	cloudtrail_createChannelCmd.MarkFlagRequired("source")
	cloudtrailCmd.AddCommand(cloudtrail_createChannelCmd)
}
