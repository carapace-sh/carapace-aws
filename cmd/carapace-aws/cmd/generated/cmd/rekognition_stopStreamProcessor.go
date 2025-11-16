package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rekognition_stopStreamProcessorCmd = &cobra.Command{
	Use:   "stop-stream-processor",
	Short: "Stops a running stream processor that was created by [CreateStreamProcessor]().",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rekognition_stopStreamProcessorCmd).Standalone()

	rekognition_stopStreamProcessorCmd.Flags().String("name", "", "The name of a stream processor created by [CreateStreamProcessor]().")
	rekognition_stopStreamProcessorCmd.MarkFlagRequired("name")
	rekognitionCmd.AddCommand(rekognition_stopStreamProcessorCmd)
}
