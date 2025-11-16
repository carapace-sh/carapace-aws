package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rekognition_startCelebrityRecognitionCmd = &cobra.Command{
	Use:   "start-celebrity-recognition",
	Short: "Starts asynchronous recognition of celebrities in a stored video.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rekognition_startCelebrityRecognitionCmd).Standalone()

	rekognition_startCelebrityRecognitionCmd.Flags().String("client-request-token", "", "Idempotent token used to identify the start request.")
	rekognition_startCelebrityRecognitionCmd.Flags().String("job-tag", "", "An identifier you specify that's returned in the completion notification that's published to your Amazon Simple Notification Service topic.")
	rekognition_startCelebrityRecognitionCmd.Flags().String("notification-channel", "", "The Amazon SNS topic ARN that you want Amazon Rekognition Video to publish the completion status of the celebrity recognition analysis to.")
	rekognition_startCelebrityRecognitionCmd.Flags().String("video", "", "The video in which you want to recognize celebrities.")
	rekognition_startCelebrityRecognitionCmd.MarkFlagRequired("video")
	rekognitionCmd.AddCommand(rekognition_startCelebrityRecognitionCmd)
}
