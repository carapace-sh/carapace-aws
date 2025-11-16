package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rekognition_startLabelDetectionCmd = &cobra.Command{
	Use:   "start-label-detection",
	Short: "Starts asynchronous detection of labels in a stored video.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rekognition_startLabelDetectionCmd).Standalone()

	rekognition_startLabelDetectionCmd.Flags().String("client-request-token", "", "Idempotent token used to identify the start request.")
	rekognition_startLabelDetectionCmd.Flags().String("features", "", "The features to return after video analysis.")
	rekognition_startLabelDetectionCmd.Flags().String("job-tag", "", "An identifier you specify that's returned in the completion notification that's published to your Amazon Simple Notification Service topic.")
	rekognition_startLabelDetectionCmd.Flags().String("min-confidence", "", "Specifies the minimum confidence that Amazon Rekognition Video must have in order to return a detected label.")
	rekognition_startLabelDetectionCmd.Flags().String("notification-channel", "", "The Amazon SNS topic ARN you want Amazon Rekognition Video to publish the completion status of the label detection operation to.")
	rekognition_startLabelDetectionCmd.Flags().String("settings", "", "The settings for a StartLabelDetection request.Contains the specified parameters for the label detection request of an asynchronous label analysis operation.")
	rekognition_startLabelDetectionCmd.Flags().String("video", "", "The video in which you want to detect labels.")
	rekognition_startLabelDetectionCmd.MarkFlagRequired("video")
	rekognitionCmd.AddCommand(rekognition_startLabelDetectionCmd)
}
