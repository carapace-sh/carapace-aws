package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_updateChannelCmd = &cobra.Command{
	Use:   "update-channel",
	Short: "Updates a channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_updateChannelCmd).Standalone()

	medialive_updateChannelCmd.Flags().String("anywhere-settings", "", "The Elemental Anywhere settings for this channel.")
	medialive_updateChannelCmd.Flags().String("cdi-input-specification", "", "Specification of CDI inputs for this channel")
	medialive_updateChannelCmd.Flags().String("channel-engine-version", "", "Channel engine version for this channel")
	medialive_updateChannelCmd.Flags().String("channel-id", "", "channel ID")
	medialive_updateChannelCmd.Flags().String("destinations", "", "A list of output destinations for this channel.")
	medialive_updateChannelCmd.Flags().String("dry-run", "", "")
	medialive_updateChannelCmd.Flags().String("encoder-settings", "", "The encoder settings for this channel.")
	medialive_updateChannelCmd.Flags().String("input-attachments", "", "")
	medialive_updateChannelCmd.Flags().String("input-specification", "", "Specification of network and file inputs for this channel")
	medialive_updateChannelCmd.Flags().String("log-level", "", "The log level to write to CloudWatch Logs.")
	medialive_updateChannelCmd.Flags().String("maintenance", "", "Maintenance settings for this channel.")
	medialive_updateChannelCmd.Flags().String("name", "", "The name of the channel.")
	medialive_updateChannelCmd.Flags().String("role-arn", "", "An optional Amazon Resource Name (ARN) of the role to assume when running the Channel.")
	medialive_updateChannelCmd.MarkFlagRequired("channel-id")
	medialiveCmd.AddCommand(medialive_updateChannelCmd)
}
