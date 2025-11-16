package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transcribe_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes the specified tags from the specified Amazon Transcribe resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transcribe_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(transcribe_untagResourceCmd).Standalone()

		transcribe_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the Amazon Transcribe resource you want to remove tags from.")
		transcribe_untagResourceCmd.Flags().String("tag-keys", "", "Removes the specified tag keys from the specified Amazon Transcribe resource.")
		transcribe_untagResourceCmd.MarkFlagRequired("resource-arn")
		transcribe_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	transcribeCmd.AddCommand(transcribe_untagResourceCmd)
}
