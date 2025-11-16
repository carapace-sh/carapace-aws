package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sqs_removePermissionCmd = &cobra.Command{
	Use:   "remove-permission",
	Short: "Revokes any permissions in the queue policy that matches the specified `Label` parameter.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sqs_removePermissionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sqs_removePermissionCmd).Standalone()

		sqs_removePermissionCmd.Flags().String("label", "", "The identification of the permission to remove.")
		sqs_removePermissionCmd.Flags().String("queue-url", "", "The URL of the Amazon SQS queue from which permissions are removed.")
		sqs_removePermissionCmd.MarkFlagRequired("label")
		sqs_removePermissionCmd.MarkFlagRequired("queue-url")
	})
	sqsCmd.AddCommand(sqs_removePermissionCmd)
}
