package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transcribe_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists all tags associated with the specified transcription job, vocabulary, model, or resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transcribe_listTagsForResourceCmd).Standalone()

	transcribe_listTagsForResourceCmd.Flags().String("resource-arn", "", "Returns a list of all tags associated with the specified Amazon Resource Name (ARN).")
	transcribe_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	transcribeCmd.AddCommand(transcribe_listTagsForResourceCmd)
}
