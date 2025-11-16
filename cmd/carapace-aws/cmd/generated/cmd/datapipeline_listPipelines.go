package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datapipeline_listPipelinesCmd = &cobra.Command{
	Use:   "list-pipelines",
	Short: "Lists the pipeline identifiers for all active pipelines that you have permission to access.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datapipeline_listPipelinesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datapipeline_listPipelinesCmd).Standalone()

		datapipeline_listPipelinesCmd.Flags().String("marker", "", "The starting point for the results to be returned.")
	})
	datapipelineCmd.AddCommand(datapipeline_listPipelinesCmd)
}
