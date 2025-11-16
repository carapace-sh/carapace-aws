package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotanalytics_updateChannelCmd = &cobra.Command{
	Use:   "update-channel",
	Short: "Used to update the settings of a channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotanalytics_updateChannelCmd).Standalone()

	iotanalytics_updateChannelCmd.Flags().String("channel-name", "", "The name of the channel to be updated.")
	iotanalytics_updateChannelCmd.Flags().String("channel-storage", "", "Where channel data is stored.")
	iotanalytics_updateChannelCmd.Flags().String("retention-period", "", "How long, in days, message data is kept for the channel.")
	iotanalytics_updateChannelCmd.MarkFlagRequired("channel-name")
	iotanalyticsCmd.AddCommand(iotanalytics_updateChannelCmd)
}
