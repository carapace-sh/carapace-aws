package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rekognition_startStreamProcessorCmd = &cobra.Command{
	Use:   "start-stream-processor",
	Short: "Starts processing a stream processor.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rekognition_startStreamProcessorCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rekognition_startStreamProcessorCmd).Standalone()

		rekognition_startStreamProcessorCmd.Flags().String("name", "", "The name of the stream processor to start processing.")
		rekognition_startStreamProcessorCmd.Flags().String("start-selector", "", "Specifies the starting point in the Kinesis stream to start processing.")
		rekognition_startStreamProcessorCmd.Flags().String("stop-selector", "", "Specifies when to stop processing the stream.")
		rekognition_startStreamProcessorCmd.MarkFlagRequired("name")
	})
	rekognitionCmd.AddCommand(rekognition_startStreamProcessorCmd)
}
