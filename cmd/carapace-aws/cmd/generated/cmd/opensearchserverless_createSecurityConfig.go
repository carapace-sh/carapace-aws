package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearchserverless_createSecurityConfigCmd = &cobra.Command{
	Use:   "create-security-config",
	Short: "Specifies a security configuration for OpenSearch Serverless.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearchserverless_createSecurityConfigCmd).Standalone()

	opensearchserverless_createSecurityConfigCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier to ensure idempotency of the request.")
	opensearchserverless_createSecurityConfigCmd.Flags().String("description", "", "A description of the security configuration.")
	opensearchserverless_createSecurityConfigCmd.Flags().String("iam-federation-options", "", "Describes IAM federation options in the form of a key-value map.")
	opensearchserverless_createSecurityConfigCmd.Flags().String("iam-identity-center-options", "", "Describes IAM Identity Center options in the form of a key-value map.")
	opensearchserverless_createSecurityConfigCmd.Flags().String("name", "", "The name of the security configuration.")
	opensearchserverless_createSecurityConfigCmd.Flags().String("saml-options", "", "Describes SAML options in in the form of a key-value map.")
	opensearchserverless_createSecurityConfigCmd.Flags().String("type", "", "The type of security configuration.")
	opensearchserverless_createSecurityConfigCmd.MarkFlagRequired("name")
	opensearchserverless_createSecurityConfigCmd.MarkFlagRequired("type")
	opensearchserverlessCmd.AddCommand(opensearchserverless_createSecurityConfigCmd)
}
