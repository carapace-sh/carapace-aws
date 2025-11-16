package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes the association of the specified tags from a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpointSmsVoiceV2_untagResourceCmd).Standalone()

		pinpointSmsVoiceV2_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
		pinpointSmsVoiceV2_untagResourceCmd.Flags().String("tag-keys", "", "An array of tag key values to unassociate with the resource.")
		pinpointSmsVoiceV2_untagResourceCmd.MarkFlagRequired("resource-arn")
		pinpointSmsVoiceV2_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_untagResourceCmd)
}
