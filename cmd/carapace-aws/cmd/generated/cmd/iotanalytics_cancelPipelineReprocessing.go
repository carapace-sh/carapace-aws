package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotanalytics_cancelPipelineReprocessingCmd = &cobra.Command{
	Use:   "cancel-pipeline-reprocessing",
	Short: "Cancels the reprocessing of data through the pipeline.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotanalytics_cancelPipelineReprocessingCmd).Standalone()

	iotanalytics_cancelPipelineReprocessingCmd.Flags().String("pipeline-name", "", "The name of pipeline for which data reprocessing is canceled.")
	iotanalytics_cancelPipelineReprocessingCmd.Flags().String("reprocessing-id", "", "The ID of the reprocessing task (returned by `StartPipelineReprocessing`).")
	iotanalytics_cancelPipelineReprocessingCmd.MarkFlagRequired("pipeline-name")
	iotanalytics_cancelPipelineReprocessingCmd.MarkFlagRequired("reprocessing-id")
	iotanalyticsCmd.AddCommand(iotanalytics_cancelPipelineReprocessingCmd)
}
