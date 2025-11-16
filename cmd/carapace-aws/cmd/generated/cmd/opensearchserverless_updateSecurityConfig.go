package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearchserverless_updateSecurityConfigCmd = &cobra.Command{
	Use:   "update-security-config",
	Short: "Updates a security configuration for OpenSearch Serverless.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearchserverless_updateSecurityConfigCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(opensearchserverless_updateSecurityConfigCmd).Standalone()

		opensearchserverless_updateSecurityConfigCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier to ensure idempotency of the request.")
		opensearchserverless_updateSecurityConfigCmd.Flags().String("config-version", "", "The version of the security configuration to be updated.")
		opensearchserverless_updateSecurityConfigCmd.Flags().String("description", "", "A description of the security configuration.")
		opensearchserverless_updateSecurityConfigCmd.Flags().String("iam-federation-options", "", "Describes IAM federation options in the form of a key-value map for updating an existing security configuration.")
		opensearchserverless_updateSecurityConfigCmd.Flags().String("iam-identity-center-options-updates", "", "Describes IAM Identity Center options in the form of a key-value map.")
		opensearchserverless_updateSecurityConfigCmd.Flags().String("id", "", "The security configuration identifier.")
		opensearchserverless_updateSecurityConfigCmd.Flags().String("saml-options", "", "SAML options in in the form of a key-value map.")
		opensearchserverless_updateSecurityConfigCmd.MarkFlagRequired("config-version")
		opensearchserverless_updateSecurityConfigCmd.MarkFlagRequired("id")
	})
	opensearchserverlessCmd.AddCommand(opensearchserverless_updateSecurityConfigCmd)
}
