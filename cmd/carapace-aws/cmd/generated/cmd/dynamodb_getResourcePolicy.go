package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dynamodb_getResourcePolicyCmd = &cobra.Command{
	Use:   "get-resource-policy",
	Short: "Returns the resource-based policy document attached to the resource, which can be a table or stream, in JSON format.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dynamodb_getResourcePolicyCmd).Standalone()

	dynamodb_getResourcePolicyCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the DynamoDB resource to which the policy is attached.")
	dynamodb_getResourcePolicyCmd.MarkFlagRequired("resource-arn")
	dynamodbCmd.AddCommand(dynamodb_getResourcePolicyCmd)
}
