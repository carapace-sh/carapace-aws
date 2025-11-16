package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotanalytics_createPipelineCmd = &cobra.Command{
	Use:   "create-pipeline",
	Short: "Creates a pipeline.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotanalytics_createPipelineCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotanalytics_createPipelineCmd).Standalone()

		iotanalytics_createPipelineCmd.Flags().String("pipeline-activities", "", "A list of `PipelineActivity` objects.")
		iotanalytics_createPipelineCmd.Flags().String("pipeline-name", "", "The name of the pipeline.")
		iotanalytics_createPipelineCmd.Flags().String("tags", "", "Metadata which can be used to manage the pipeline.")
		iotanalytics_createPipelineCmd.MarkFlagRequired("pipeline-activities")
		iotanalytics_createPipelineCmd.MarkFlagRequired("pipeline-name")
	})
	iotanalyticsCmd.AddCommand(iotanalytics_createPipelineCmd)
}
