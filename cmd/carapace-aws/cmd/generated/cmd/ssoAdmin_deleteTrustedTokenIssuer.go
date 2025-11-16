package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssoAdmin_deleteTrustedTokenIssuerCmd = &cobra.Command{
	Use:   "delete-trusted-token-issuer",
	Short: "Deletes a trusted token issuer configuration from an instance of IAM Identity Center.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssoAdmin_deleteTrustedTokenIssuerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssoAdmin_deleteTrustedTokenIssuerCmd).Standalone()

		ssoAdmin_deleteTrustedTokenIssuerCmd.Flags().String("trusted-token-issuer-arn", "", "Specifies the ARN of the trusted token issuer configuration to delete.")
		ssoAdmin_deleteTrustedTokenIssuerCmd.MarkFlagRequired("trusted-token-issuer-arn")
	})
	ssoAdminCmd.AddCommand(ssoAdmin_deleteTrustedTokenIssuerCmd)
}
