package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datapipeline_createPipelineCmd = &cobra.Command{
	Use:   "create-pipeline",
	Short: "Creates a new, empty pipeline.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datapipeline_createPipelineCmd).Standalone()

	datapipeline_createPipelineCmd.Flags().String("description", "", "The description for the pipeline.")
	datapipeline_createPipelineCmd.Flags().String("name", "", "The name for the pipeline.")
	datapipeline_createPipelineCmd.Flags().String("tags", "", "A list of tags to associate with the pipeline at creation.")
	datapipeline_createPipelineCmd.Flags().String("unique-id", "", "A unique identifier.")
	datapipeline_createPipelineCmd.MarkFlagRequired("name")
	datapipeline_createPipelineCmd.MarkFlagRequired("unique-id")
	datapipelineCmd.AddCommand(datapipeline_createPipelineCmd)
}
