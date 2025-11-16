package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotanalytics_updatePipelineCmd = &cobra.Command{
	Use:   "update-pipeline",
	Short: "Updates the settings of a pipeline.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotanalytics_updatePipelineCmd).Standalone()

	iotanalytics_updatePipelineCmd.Flags().String("pipeline-activities", "", "A list of `PipelineActivity` objects.")
	iotanalytics_updatePipelineCmd.Flags().String("pipeline-name", "", "The name of the pipeline to update.")
	iotanalytics_updatePipelineCmd.MarkFlagRequired("pipeline-activities")
	iotanalytics_updatePipelineCmd.MarkFlagRequired("pipeline-name")
	iotanalyticsCmd.AddCommand(iotanalytics_updatePipelineCmd)
}
