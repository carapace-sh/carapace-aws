package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datapipeline_describeObjectsCmd = &cobra.Command{
	Use:   "describe-objects",
	Short: "Gets the object definitions for a set of objects associated with the pipeline.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datapipeline_describeObjectsCmd).Standalone()

	datapipeline_describeObjectsCmd.Flags().String("evaluate-expressions", "", "Indicates whether any expressions in the object should be evaluated when the object descriptions are returned.")
	datapipeline_describeObjectsCmd.Flags().String("marker", "", "The starting point for the results to be returned.")
	datapipeline_describeObjectsCmd.Flags().String("object-ids", "", "The IDs of the pipeline objects that contain the definitions to be described.")
	datapipeline_describeObjectsCmd.Flags().String("pipeline-id", "", "The ID of the pipeline that contains the object definitions.")
	datapipeline_describeObjectsCmd.MarkFlagRequired("object-ids")
	datapipeline_describeObjectsCmd.MarkFlagRequired("pipeline-id")
	datapipelineCmd.AddCommand(datapipeline_describeObjectsCmd)
}
