package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datapipeline_addTagsCmd = &cobra.Command{
	Use:   "add-tags",
	Short: "Adds or modifies tags for the specified pipeline.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datapipeline_addTagsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datapipeline_addTagsCmd).Standalone()

		datapipeline_addTagsCmd.Flags().String("pipeline-id", "", "The ID of the pipeline.")
		datapipeline_addTagsCmd.Flags().String("tags", "", "The tags to add, as key/value pairs.")
		datapipeline_addTagsCmd.MarkFlagRequired("pipeline-id")
		datapipeline_addTagsCmd.MarkFlagRequired("tags")
	})
	datapipelineCmd.AddCommand(datapipeline_addTagsCmd)
}
