package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediapackagev2_getChannelPolicyCmd = &cobra.Command{
	Use:   "get-channel-policy",
	Short: "Retrieves the specified channel policy that's configured in AWS Elemental MediaPackage.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediapackagev2_getChannelPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediapackagev2_getChannelPolicyCmd).Standalone()

		mediapackagev2_getChannelPolicyCmd.Flags().String("channel-group-name", "", "The name that describes the channel group.")
		mediapackagev2_getChannelPolicyCmd.Flags().String("channel-name", "", "The name that describes the channel.")
		mediapackagev2_getChannelPolicyCmd.MarkFlagRequired("channel-group-name")
		mediapackagev2_getChannelPolicyCmd.MarkFlagRequired("channel-name")
	})
	mediapackagev2Cmd.AddCommand(mediapackagev2_getChannelPolicyCmd)
}
