package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotanalytics_deleteChannelCmd = &cobra.Command{
	Use:   "delete-channel",
	Short: "Deletes the specified channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotanalytics_deleteChannelCmd).Standalone()

	iotanalytics_deleteChannelCmd.Flags().String("channel-name", "", "The name of the channel to delete.")
	iotanalytics_deleteChannelCmd.MarkFlagRequired("channel-name")
	iotanalyticsCmd.AddCommand(iotanalytics_deleteChannelCmd)
}
