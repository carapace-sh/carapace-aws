package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rekognition_startSegmentDetectionCmd = &cobra.Command{
	Use:   "start-segment-detection",
	Short: "Starts asynchronous detection of segment detection in a stored video.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rekognition_startSegmentDetectionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rekognition_startSegmentDetectionCmd).Standalone()

		rekognition_startSegmentDetectionCmd.Flags().String("client-request-token", "", "Idempotent token used to identify the start request.")
		rekognition_startSegmentDetectionCmd.Flags().String("filters", "", "Filters for technical cue or shot detection.")
		rekognition_startSegmentDetectionCmd.Flags().String("job-tag", "", "An identifier you specify that's returned in the completion notification that's published to your Amazon Simple Notification Service topic.")
		rekognition_startSegmentDetectionCmd.Flags().String("notification-channel", "", "The ARN of the Amazon SNS topic to which you want Amazon Rekognition Video to publish the completion status of the segment detection operation.")
		rekognition_startSegmentDetectionCmd.Flags().String("segment-types", "", "An array of segment types to detect in the video.")
		rekognition_startSegmentDetectionCmd.Flags().String("video", "", "")
		rekognition_startSegmentDetectionCmd.MarkFlagRequired("segment-types")
		rekognition_startSegmentDetectionCmd.MarkFlagRequired("video")
	})
	rekognitionCmd.AddCommand(rekognition_startSegmentDetectionCmd)
}
