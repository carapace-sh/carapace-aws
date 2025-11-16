package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var voiceId_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Tags a Voice ID resource with the provided list of tags.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(voiceId_tagResourceCmd).Standalone()

	voiceId_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the Voice ID resource you want to tag.")
	voiceId_tagResourceCmd.Flags().String("tags", "", "The list of tags to assign to the specified resource.")
	voiceId_tagResourceCmd.MarkFlagRequired("resource-arn")
	voiceId_tagResourceCmd.MarkFlagRequired("tags")
	voiceIdCmd.AddCommand(voiceId_tagResourceCmd)
}
