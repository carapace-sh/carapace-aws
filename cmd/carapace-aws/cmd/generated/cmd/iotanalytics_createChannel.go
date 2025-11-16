package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotanalytics_createChannelCmd = &cobra.Command{
	Use:   "create-channel",
	Short: "Used to create a channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotanalytics_createChannelCmd).Standalone()

	iotanalytics_createChannelCmd.Flags().String("channel-name", "", "The name of the channel.")
	iotanalytics_createChannelCmd.Flags().String("channel-storage", "", "Where channel data is stored.")
	iotanalytics_createChannelCmd.Flags().String("retention-period", "", "How long, in days, message data is kept for the channel.")
	iotanalytics_createChannelCmd.Flags().String("tags", "", "Metadata which can be used to manage the channel.")
	iotanalytics_createChannelCmd.MarkFlagRequired("channel-name")
	iotanalyticsCmd.AddCommand(iotanalytics_createChannelCmd)
}
