package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds or overwrites only the specified tags for the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_tagResourceCmd).Standalone()

	pinpointSmsVoiceV2_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
	pinpointSmsVoiceV2_tagResourceCmd.Flags().String("tags", "", "An array of key and value pair tags that are associated with the resource.")
	pinpointSmsVoiceV2_tagResourceCmd.MarkFlagRequired("resource-arn")
	pinpointSmsVoiceV2_tagResourceCmd.MarkFlagRequired("tags")
	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_tagResourceCmd)
}
