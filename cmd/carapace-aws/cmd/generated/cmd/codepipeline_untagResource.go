package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codepipeline_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes tags from an Amazon Web Services resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codepipeline_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codepipeline_untagResourceCmd).Standalone()

		codepipeline_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource to remove tags from.")
		codepipeline_untagResourceCmd.Flags().String("tag-keys", "", "The list of keys for the tags to be removed from the resource.")
		codepipeline_untagResourceCmd.MarkFlagRequired("resource-arn")
		codepipeline_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	codepipelineCmd.AddCommand(codepipeline_untagResourceCmd)
}
