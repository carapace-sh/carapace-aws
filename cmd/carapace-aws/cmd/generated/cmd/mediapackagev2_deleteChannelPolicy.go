package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediapackagev2_deleteChannelPolicyCmd = &cobra.Command{
	Use:   "delete-channel-policy",
	Short: "Delete a channel policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediapackagev2_deleteChannelPolicyCmd).Standalone()

	mediapackagev2_deleteChannelPolicyCmd.Flags().String("channel-group-name", "", "The name that describes the channel group.")
	mediapackagev2_deleteChannelPolicyCmd.Flags().String("channel-name", "", "The name that describes the channel.")
	mediapackagev2_deleteChannelPolicyCmd.MarkFlagRequired("channel-group-name")
	mediapackagev2_deleteChannelPolicyCmd.MarkFlagRequired("channel-name")
	mediapackagev2Cmd.AddCommand(mediapackagev2_deleteChannelPolicyCmd)
}
