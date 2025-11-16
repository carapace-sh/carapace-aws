package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rekognition_startFaceSearchCmd = &cobra.Command{
	Use:   "start-face-search",
	Short: "Starts the asynchronous search for faces in a collection that match the faces of persons detected in a stored video.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rekognition_startFaceSearchCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rekognition_startFaceSearchCmd).Standalone()

		rekognition_startFaceSearchCmd.Flags().String("client-request-token", "", "Idempotent token used to identify the start request.")
		rekognition_startFaceSearchCmd.Flags().String("collection-id", "", "ID of the collection that contains the faces you want to search for.")
		rekognition_startFaceSearchCmd.Flags().String("face-match-threshold", "", "The minimum confidence in the person match to return.")
		rekognition_startFaceSearchCmd.Flags().String("job-tag", "", "An identifier you specify that's returned in the completion notification that's published to your Amazon Simple Notification Service topic.")
		rekognition_startFaceSearchCmd.Flags().String("notification-channel", "", "The ARN of the Amazon SNS topic to which you want Amazon Rekognition Video to publish the completion status of the search.")
		rekognition_startFaceSearchCmd.Flags().String("video", "", "The video you want to search.")
		rekognition_startFaceSearchCmd.MarkFlagRequired("collection-id")
		rekognition_startFaceSearchCmd.MarkFlagRequired("video")
	})
	rekognitionCmd.AddCommand(rekognition_startFaceSearchCmd)
}
