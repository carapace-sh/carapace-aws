package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sqs_addPermissionCmd = &cobra.Command{
	Use:   "add-permission",
	Short: "Adds a permission to a queue for a specific [principal](https://docs.aws.amazon.com/general/latest/gr/glos-chap.html#P).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sqs_addPermissionCmd).Standalone()

	sqs_addPermissionCmd.Flags().String("actions", "", "The action the client wants to allow for the specified principal.")
	sqs_addPermissionCmd.Flags().String("awsaccount-ids", "", "The Amazon Web Services account numbers of the [principals](https://docs.aws.amazon.com/general/latest/gr/glos-chap.html#P) who are to receive permission.")
	sqs_addPermissionCmd.Flags().String("label", "", "The unique identification of the permission you're setting (for example, `AliceSendMessage`).")
	sqs_addPermissionCmd.Flags().String("queue-url", "", "The URL of the Amazon SQS queue to which permissions are added.")
	sqs_addPermissionCmd.MarkFlagRequired("actions")
	sqs_addPermissionCmd.MarkFlagRequired("awsaccount-ids")
	sqs_addPermissionCmd.MarkFlagRequired("label")
	sqs_addPermissionCmd.MarkFlagRequired("queue-url")
	sqsCmd.AddCommand(sqs_addPermissionCmd)
}
