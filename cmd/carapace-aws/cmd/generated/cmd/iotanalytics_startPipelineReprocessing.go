package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotanalytics_startPipelineReprocessingCmd = &cobra.Command{
	Use:   "start-pipeline-reprocessing",
	Short: "Starts the reprocessing of raw message data through the pipeline.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotanalytics_startPipelineReprocessingCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotanalytics_startPipelineReprocessingCmd).Standalone()

		iotanalytics_startPipelineReprocessingCmd.Flags().String("channel-messages", "", "Specifies one or more sets of channel messages that you want to reprocess.")
		iotanalytics_startPipelineReprocessingCmd.Flags().String("end-time", "", "The end time (exclusive) of raw message data that is reprocessed.")
		iotanalytics_startPipelineReprocessingCmd.Flags().String("pipeline-name", "", "The name of the pipeline on which to start reprocessing.")
		iotanalytics_startPipelineReprocessingCmd.Flags().String("start-time", "", "The start time (inclusive) of raw message data that is reprocessed.")
		iotanalytics_startPipelineReprocessingCmd.MarkFlagRequired("pipeline-name")
	})
	iotanalyticsCmd.AddCommand(iotanalytics_startPipelineReprocessingCmd)
}
