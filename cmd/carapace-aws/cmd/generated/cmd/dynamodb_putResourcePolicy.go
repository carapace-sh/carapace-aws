package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dynamodb_putResourcePolicyCmd = &cobra.Command{
	Use:   "put-resource-policy",
	Short: "Attaches a resource-based policy document to the resource, which can be a table or stream.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dynamodb_putResourcePolicyCmd).Standalone()

	dynamodb_putResourcePolicyCmd.Flags().String("confirm-remove-self-resource-access", "", "Set this parameter to `true` to confirm that you want to remove your permissions to change the policy of this resource in the future.")
	dynamodb_putResourcePolicyCmd.Flags().String("expected-revision-id", "", "A string value that you can use to conditionally update your policy.")
	dynamodb_putResourcePolicyCmd.Flags().String("policy", "", "An Amazon Web Services resource-based policy document in JSON format.")
	dynamodb_putResourcePolicyCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the DynamoDB resource to which the policy will be attached.")
	dynamodb_putResourcePolicyCmd.MarkFlagRequired("policy")
	dynamodb_putResourcePolicyCmd.MarkFlagRequired("resource-arn")
	dynamodbCmd.AddCommand(dynamodb_putResourcePolicyCmd)
}
