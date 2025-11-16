package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rekognition_startFaceDetectionCmd = &cobra.Command{
	Use:   "start-face-detection",
	Short: "Starts asynchronous detection of faces in a stored video.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rekognition_startFaceDetectionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rekognition_startFaceDetectionCmd).Standalone()

		rekognition_startFaceDetectionCmd.Flags().String("client-request-token", "", "Idempotent token used to identify the start request.")
		rekognition_startFaceDetectionCmd.Flags().String("face-attributes", "", "The face attributes you want returned.")
		rekognition_startFaceDetectionCmd.Flags().String("job-tag", "", "An identifier you specify that's returned in the completion notification that's published to your Amazon Simple Notification Service topic.")
		rekognition_startFaceDetectionCmd.Flags().String("notification-channel", "", "The ARN of the Amazon SNS topic to which you want Amazon Rekognition Video to publish the completion status of the face detection operation.")
		rekognition_startFaceDetectionCmd.Flags().String("video", "", "The video in which you want to detect faces.")
		rekognition_startFaceDetectionCmd.MarkFlagRequired("video")
	})
	rekognitionCmd.AddCommand(rekognition_startFaceDetectionCmd)
}
