package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datapipeline_removeTagsCmd = &cobra.Command{
	Use:   "remove-tags",
	Short: "Removes existing tags from the specified pipeline.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datapipeline_removeTagsCmd).Standalone()

	datapipeline_removeTagsCmd.Flags().String("pipeline-id", "", "The ID of the pipeline.")
	datapipeline_removeTagsCmd.Flags().String("tag-keys", "", "The keys of the tags to remove.")
	datapipeline_removeTagsCmd.MarkFlagRequired("pipeline-id")
	datapipeline_removeTagsCmd.MarkFlagRequired("tag-keys")
	datapipelineCmd.AddCommand(datapipeline_removeTagsCmd)
}
