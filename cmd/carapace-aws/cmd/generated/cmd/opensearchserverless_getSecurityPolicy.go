package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearchserverless_getSecurityPolicyCmd = &cobra.Command{
	Use:   "get-security-policy",
	Short: "Returns information about a configured OpenSearch Serverless security policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearchserverless_getSecurityPolicyCmd).Standalone()

	opensearchserverless_getSecurityPolicyCmd.Flags().String("name", "", "The name of the security policy.")
	opensearchserverless_getSecurityPolicyCmd.Flags().String("type", "", "The type of security policy.")
	opensearchserverless_getSecurityPolicyCmd.MarkFlagRequired("name")
	opensearchserverless_getSecurityPolicyCmd.MarkFlagRequired("type")
	opensearchserverlessCmd.AddCommand(opensearchserverless_getSecurityPolicyCmd)
}
