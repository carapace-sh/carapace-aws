package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rekognition_startTextDetectionCmd = &cobra.Command{
	Use:   "start-text-detection",
	Short: "Starts asynchronous detection of text in a stored video.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rekognition_startTextDetectionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rekognition_startTextDetectionCmd).Standalone()

		rekognition_startTextDetectionCmd.Flags().String("client-request-token", "", "Idempotent token used to identify the start request.")
		rekognition_startTextDetectionCmd.Flags().String("filters", "", "Optional parameters that let you set criteria the text must meet to be included in your response.")
		rekognition_startTextDetectionCmd.Flags().String("job-tag", "", "An identifier returned in the completion status published by your Amazon Simple Notification Service topic.")
		rekognition_startTextDetectionCmd.Flags().String("notification-channel", "", "")
		rekognition_startTextDetectionCmd.Flags().String("video", "", "")
		rekognition_startTextDetectionCmd.MarkFlagRequired("video")
	})
	rekognitionCmd.AddCommand(rekognition_startTextDetectionCmd)
}
