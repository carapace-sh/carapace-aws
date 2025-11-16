package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes tags from a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_untagResourceCmd).Standalone()

	chimeSdkVoice_untagResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource having its tags removed.")
	chimeSdkVoice_untagResourceCmd.Flags().String("tag-keys", "", "The keys of the tags being removed from the resource.")
	chimeSdkVoice_untagResourceCmd.MarkFlagRequired("resource-arn")
	chimeSdkVoice_untagResourceCmd.MarkFlagRequired("tag-keys")
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_untagResourceCmd)
}
