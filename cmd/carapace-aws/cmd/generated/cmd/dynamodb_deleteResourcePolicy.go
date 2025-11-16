package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dynamodb_deleteResourcePolicyCmd = &cobra.Command{
	Use:   "delete-resource-policy",
	Short: "Deletes the resource-based policy attached to the resource, which can be a table or stream.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dynamodb_deleteResourcePolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dynamodb_deleteResourcePolicyCmd).Standalone()

		dynamodb_deleteResourcePolicyCmd.Flags().String("expected-revision-id", "", "A string value that you can use to conditionally delete your policy.")
		dynamodb_deleteResourcePolicyCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the DynamoDB resource from which the policy will be removed.")
		dynamodb_deleteResourcePolicyCmd.MarkFlagRequired("resource-arn")
	})
	dynamodbCmd.AddCommand(dynamodb_deleteResourcePolicyCmd)
}
