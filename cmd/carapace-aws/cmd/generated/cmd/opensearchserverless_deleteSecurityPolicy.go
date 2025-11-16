package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearchserverless_deleteSecurityPolicyCmd = &cobra.Command{
	Use:   "delete-security-policy",
	Short: "Deletes an OpenSearch Serverless security policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearchserverless_deleteSecurityPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(opensearchserverless_deleteSecurityPolicyCmd).Standalone()

		opensearchserverless_deleteSecurityPolicyCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier to ensure idempotency of the request.")
		opensearchserverless_deleteSecurityPolicyCmd.Flags().String("name", "", "The name of the policy to delete.")
		opensearchserverless_deleteSecurityPolicyCmd.Flags().String("type", "", "The type of policy.")
		opensearchserverless_deleteSecurityPolicyCmd.MarkFlagRequired("name")
		opensearchserverless_deleteSecurityPolicyCmd.MarkFlagRequired("type")
	})
	opensearchserverlessCmd.AddCommand(opensearchserverless_deleteSecurityPolicyCmd)
}
