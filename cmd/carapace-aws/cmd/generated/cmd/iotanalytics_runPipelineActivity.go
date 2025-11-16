package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotanalytics_runPipelineActivityCmd = &cobra.Command{
	Use:   "run-pipeline-activity",
	Short: "Simulates the results of running a pipeline activity on a message payload.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotanalytics_runPipelineActivityCmd).Standalone()

	iotanalytics_runPipelineActivityCmd.Flags().String("payloads", "", "The sample message payloads on which the pipeline activity is run.")
	iotanalytics_runPipelineActivityCmd.Flags().String("pipeline-activity", "", "The pipeline activity that is run.")
	iotanalytics_runPipelineActivityCmd.MarkFlagRequired("payloads")
	iotanalytics_runPipelineActivityCmd.MarkFlagRequired("pipeline-activity")
	iotanalyticsCmd.AddCommand(iotanalytics_runPipelineActivityCmd)
}
