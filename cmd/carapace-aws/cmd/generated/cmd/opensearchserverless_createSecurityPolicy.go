package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearchserverless_createSecurityPolicyCmd = &cobra.Command{
	Use:   "create-security-policy",
	Short: "Creates a security policy to be used by one or more OpenSearch Serverless collections.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearchserverless_createSecurityPolicyCmd).Standalone()

	opensearchserverless_createSecurityPolicyCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier to ensure idempotency of the request.")
	opensearchserverless_createSecurityPolicyCmd.Flags().String("description", "", "A description of the policy.")
	opensearchserverless_createSecurityPolicyCmd.Flags().String("name", "", "The name of the policy.")
	opensearchserverless_createSecurityPolicyCmd.Flags().String("policy", "", "The JSON policy document to use as the content for the new policy.")
	opensearchserverless_createSecurityPolicyCmd.Flags().String("type", "", "The type of security policy.")
	opensearchserverless_createSecurityPolicyCmd.MarkFlagRequired("name")
	opensearchserverless_createSecurityPolicyCmd.MarkFlagRequired("policy")
	opensearchserverless_createSecurityPolicyCmd.MarkFlagRequired("type")
	opensearchserverlessCmd.AddCommand(opensearchserverless_createSecurityPolicyCmd)
}
