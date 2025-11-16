package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rekognition_startPersonTrackingCmd = &cobra.Command{
	Use:   "start-person-tracking",
	Short: "*End of support notice:* On October 31, 2025, AWS will discontinue support for Amazon Rekognition People Pathing.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rekognition_startPersonTrackingCmd).Standalone()

	rekognition_startPersonTrackingCmd.Flags().String("client-request-token", "", "Idempotent token used to identify the start request.")
	rekognition_startPersonTrackingCmd.Flags().String("job-tag", "", "An identifier you specify that's returned in the completion notification that's published to your Amazon Simple Notification Service topic.")
	rekognition_startPersonTrackingCmd.Flags().String("notification-channel", "", "The Amazon SNS topic ARN you want Amazon Rekognition Video to publish the completion status of the people detection operation to.")
	rekognition_startPersonTrackingCmd.Flags().String("video", "", "The video in which you want to detect people.")
	rekognition_startPersonTrackingCmd.MarkFlagRequired("video")
	rekognitionCmd.AddCommand(rekognition_startPersonTrackingCmd)
}
