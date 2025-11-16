package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sns_publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "Sends a message to an Amazon SNS topic, a text message (SMS message) directly to a phone number, or a message to a mobile platform endpoint (when you specify the `TargetArn`).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sns_publishCmd).Standalone()

	sns_publishCmd.Flags().String("message", "", "The message you want to send.")
	sns_publishCmd.Flags().String("message-attributes", "", "Message attributes for Publish action.")
	sns_publishCmd.Flags().String("message-deduplication-id", "", "- This parameter applies only to FIFO (first-in-first-out) topics.")
	sns_publishCmd.Flags().String("message-group-id", "", "The `MessageGroupId` can contain up to 128 alphanumeric characters `(a-z, A-Z, 0-9)` and punctuation ``(!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~)``.")
	sns_publishCmd.Flags().String("message-structure", "", "Set `MessageStructure` to `json` if you want to send a different message for each protocol.")
	sns_publishCmd.Flags().String("phone-number", "", "The phone number to which you want to deliver an SMS message.")
	sns_publishCmd.Flags().String("subject", "", "Optional parameter to be used as the \"Subject\" line when the message is delivered to email endpoints.")
	sns_publishCmd.Flags().String("target-arn", "", "If you don't specify a value for the `TargetArn` parameter, you must specify a value for the `PhoneNumber` or `TopicArn` parameters.")
	sns_publishCmd.Flags().String("topic-arn", "", "The topic you want to publish to.")
	sns_publishCmd.MarkFlagRequired("message")
	snsCmd.AddCommand(sns_publishCmd)
}
