package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datapipeline_describePipelinesCmd = &cobra.Command{
	Use:   "describe-pipelines",
	Short: "Retrieves metadata about one or more pipelines.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datapipeline_describePipelinesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datapipeline_describePipelinesCmd).Standalone()

		datapipeline_describePipelinesCmd.Flags().String("pipeline-ids", "", "The IDs of the pipelines to describe.")
		datapipeline_describePipelinesCmd.MarkFlagRequired("pipeline-ids")
	})
	datapipelineCmd.AddCommand(datapipeline_describePipelinesCmd)
}
