package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssoAdmin_describeTrustedTokenIssuerCmd = &cobra.Command{
	Use:   "describe-trusted-token-issuer",
	Short: "Retrieves details about a trusted token issuer configuration stored in an instance of IAM Identity Center.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssoAdmin_describeTrustedTokenIssuerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssoAdmin_describeTrustedTokenIssuerCmd).Standalone()

		ssoAdmin_describeTrustedTokenIssuerCmd.Flags().String("trusted-token-issuer-arn", "", "Specifies the ARN of the trusted token issuer configuration that you want details about.")
		ssoAdmin_describeTrustedTokenIssuerCmd.MarkFlagRequired("trusted-token-issuer-arn")
	})
	ssoAdminCmd.AddCommand(ssoAdmin_describeTrustedTokenIssuerCmd)
}
