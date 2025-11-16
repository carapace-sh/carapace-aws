package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codepipeline_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds to or modifies the tags of the given resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codepipeline_tagResourceCmd).Standalone()

	codepipeline_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource you want to add tags to.")
	codepipeline_tagResourceCmd.Flags().String("tags", "", "The tags you want to modify or add to the resource.")
	codepipeline_tagResourceCmd.MarkFlagRequired("resource-arn")
	codepipeline_tagResourceCmd.MarkFlagRequired("tags")
	codepipelineCmd.AddCommand(codepipeline_tagResourceCmd)
}
