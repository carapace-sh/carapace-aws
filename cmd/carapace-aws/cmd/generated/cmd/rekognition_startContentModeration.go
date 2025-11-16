package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rekognition_startContentModerationCmd = &cobra.Command{
	Use:   "start-content-moderation",
	Short: "Starts asynchronous detection of inappropriate, unwanted, or offensive content in a stored video.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rekognition_startContentModerationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rekognition_startContentModerationCmd).Standalone()

		rekognition_startContentModerationCmd.Flags().String("client-request-token", "", "Idempotent token used to identify the start request.")
		rekognition_startContentModerationCmd.Flags().String("job-tag", "", "An identifier you specify that's returned in the completion notification that's published to your Amazon Simple Notification Service topic.")
		rekognition_startContentModerationCmd.Flags().String("min-confidence", "", "Specifies the minimum confidence that Amazon Rekognition must have in order to return a moderated content label.")
		rekognition_startContentModerationCmd.Flags().String("notification-channel", "", "The Amazon SNS topic ARN that you want Amazon Rekognition Video to publish the completion status of the content analysis to.")
		rekognition_startContentModerationCmd.Flags().String("video", "", "The video in which you want to detect inappropriate, unwanted, or offensive content.")
		rekognition_startContentModerationCmd.MarkFlagRequired("video")
	})
	rekognitionCmd.AddCommand(rekognition_startContentModerationCmd)
}
