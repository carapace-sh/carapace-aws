package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var voiceId_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists all tags associated with a specified Voice ID resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(voiceId_listTagsForResourceCmd).Standalone()

	voiceId_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the Voice ID resource for which you want to list the tags.")
	voiceId_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	voiceIdCmd.AddCommand(voiceId_listTagsForResourceCmd)
}
