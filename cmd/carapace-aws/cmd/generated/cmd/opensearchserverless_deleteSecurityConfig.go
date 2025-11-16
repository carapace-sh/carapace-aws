package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearchserverless_deleteSecurityConfigCmd = &cobra.Command{
	Use:   "delete-security-config",
	Short: "Deletes a security configuration for OpenSearch Serverless.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearchserverless_deleteSecurityConfigCmd).Standalone()

	opensearchserverless_deleteSecurityConfigCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier to ensure idempotency of the request.")
	opensearchserverless_deleteSecurityConfigCmd.Flags().String("id", "", "The security configuration identifier.")
	opensearchserverless_deleteSecurityConfigCmd.MarkFlagRequired("id")
	opensearchserverlessCmd.AddCommand(opensearchserverless_deleteSecurityConfigCmd)
}
