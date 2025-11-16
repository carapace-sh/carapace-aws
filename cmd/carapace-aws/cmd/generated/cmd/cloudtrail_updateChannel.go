package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudtrail_updateChannelCmd = &cobra.Command{
	Use:   "update-channel",
	Short: "Updates a channel specified by a required channel ARN or UUID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudtrail_updateChannelCmd).Standalone()

	cloudtrail_updateChannelCmd.Flags().String("channel", "", "The ARN or ID (the ARN suffix) of the channel that you want to update.")
	cloudtrail_updateChannelCmd.Flags().String("destinations", "", "The ARNs of event data stores that you want to log events arriving through the channel.")
	cloudtrail_updateChannelCmd.Flags().String("name", "", "Changes the name of the channel.")
	cloudtrail_updateChannelCmd.MarkFlagRequired("channel")
	cloudtrailCmd.AddCommand(cloudtrail_updateChannelCmd)
}
