package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearchserverless_deleteAccessPolicyCmd = &cobra.Command{
	Use:   "delete-access-policy",
	Short: "Deletes an OpenSearch Serverless access policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearchserverless_deleteAccessPolicyCmd).Standalone()

	opensearchserverless_deleteAccessPolicyCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier to ensure idempotency of the request.")
	opensearchserverless_deleteAccessPolicyCmd.Flags().String("name", "", "The name of the policy to delete.")
	opensearchserverless_deleteAccessPolicyCmd.Flags().String("type", "", "The type of policy.")
	opensearchserverless_deleteAccessPolicyCmd.MarkFlagRequired("name")
	opensearchserverless_deleteAccessPolicyCmd.MarkFlagRequired("type")
	opensearchserverlessCmd.AddCommand(opensearchserverless_deleteAccessPolicyCmd)
}
