package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datapipeline_queryObjectsCmd = &cobra.Command{
	Use:   "query-objects",
	Short: "Queries the specified pipeline for the names of objects that match the specified set of conditions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datapipeline_queryObjectsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datapipeline_queryObjectsCmd).Standalone()

		datapipeline_queryObjectsCmd.Flags().String("limit", "", "The maximum number of object names that `QueryObjects` will return in a single call.")
		datapipeline_queryObjectsCmd.Flags().String("marker", "", "The starting point for the results to be returned.")
		datapipeline_queryObjectsCmd.Flags().String("pipeline-id", "", "The ID of the pipeline.")
		datapipeline_queryObjectsCmd.Flags().String("query", "", "The query that defines the objects to be returned.")
		datapipeline_queryObjectsCmd.Flags().String("sphere", "", "Indicates whether the query applies to components or instances.")
		datapipeline_queryObjectsCmd.MarkFlagRequired("pipeline-id")
		datapipeline_queryObjectsCmd.MarkFlagRequired("sphere")
	})
	datapipelineCmd.AddCommand(datapipeline_queryObjectsCmd)
}
