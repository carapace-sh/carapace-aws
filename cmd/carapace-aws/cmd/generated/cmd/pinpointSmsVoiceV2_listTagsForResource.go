package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "List all tags associated with a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpointSmsVoiceV2_listTagsForResourceCmd).Standalone()

		pinpointSmsVoiceV2_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource to query for.")
		pinpointSmsVoiceV2_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_listTagsForResourceCmd)
}
