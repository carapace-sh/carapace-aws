package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rekognition_describeStreamProcessorCmd = &cobra.Command{
	Use:   "describe-stream-processor",
	Short: "Provides information about a stream processor created by [CreateStreamProcessor]().",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rekognition_describeStreamProcessorCmd).Standalone()

	rekognition_describeStreamProcessorCmd.Flags().String("name", "", "Name of the stream processor for which you want information.")
	rekognition_describeStreamProcessorCmd.MarkFlagRequired("name")
	rekognitionCmd.AddCommand(rekognition_describeStreamProcessorCmd)
}
