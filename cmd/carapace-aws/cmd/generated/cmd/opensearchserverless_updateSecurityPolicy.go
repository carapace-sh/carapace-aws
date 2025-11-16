package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearchserverless_updateSecurityPolicyCmd = &cobra.Command{
	Use:   "update-security-policy",
	Short: "Updates an OpenSearch Serverless security policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearchserverless_updateSecurityPolicyCmd).Standalone()

	opensearchserverless_updateSecurityPolicyCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier to ensure idempotency of the request.")
	opensearchserverless_updateSecurityPolicyCmd.Flags().String("description", "", "A description of the policy.")
	opensearchserverless_updateSecurityPolicyCmd.Flags().String("name", "", "The name of the policy.")
	opensearchserverless_updateSecurityPolicyCmd.Flags().String("policy", "", "The JSON policy document to use as the content for the new policy.")
	opensearchserverless_updateSecurityPolicyCmd.Flags().String("policy-version", "", "The version of the policy being updated.")
	opensearchserverless_updateSecurityPolicyCmd.Flags().String("type", "", "The type of access policy.")
	opensearchserverless_updateSecurityPolicyCmd.MarkFlagRequired("name")
	opensearchserverless_updateSecurityPolicyCmd.MarkFlagRequired("policy-version")
	opensearchserverless_updateSecurityPolicyCmd.MarkFlagRequired("type")
	opensearchserverlessCmd.AddCommand(opensearchserverless_updateSecurityPolicyCmd)
}
