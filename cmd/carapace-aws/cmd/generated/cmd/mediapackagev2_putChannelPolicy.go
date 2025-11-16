package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediapackagev2_putChannelPolicyCmd = &cobra.Command{
	Use:   "put-channel-policy",
	Short: "Attaches an IAM policy to the specified channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediapackagev2_putChannelPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediapackagev2_putChannelPolicyCmd).Standalone()

		mediapackagev2_putChannelPolicyCmd.Flags().String("channel-group-name", "", "The name that describes the channel group.")
		mediapackagev2_putChannelPolicyCmd.Flags().String("channel-name", "", "The name that describes the channel.")
		mediapackagev2_putChannelPolicyCmd.Flags().String("policy", "", "The policy to attach to the specified channel.")
		mediapackagev2_putChannelPolicyCmd.MarkFlagRequired("channel-group-name")
		mediapackagev2_putChannelPolicyCmd.MarkFlagRequired("channel-name")
		mediapackagev2_putChannelPolicyCmd.MarkFlagRequired("policy")
	})
	mediapackagev2Cmd.AddCommand(mediapackagev2_putChannelPolicyCmd)
}
