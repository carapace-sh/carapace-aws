package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var voiceId_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes specified tags from a specified Amazon Connect Voice ID resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(voiceId_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(voiceId_untagResourceCmd).Standalone()

		voiceId_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the Voice ID resource you want to remove tags from.")
		voiceId_untagResourceCmd.Flags().String("tag-keys", "", "The list of tag keys you want to remove from the specified resource.")
		voiceId_untagResourceCmd.MarkFlagRequired("resource-arn")
		voiceId_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	voiceIdCmd.AddCommand(voiceId_untagResourceCmd)
}
