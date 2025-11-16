package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssoAdmin_updateTrustedTokenIssuerCmd = &cobra.Command{
	Use:   "update-trusted-token-issuer",
	Short: "Updates the name of the trusted token issuer, or the path of a source attribute or destination attribute for a trusted token issuer configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssoAdmin_updateTrustedTokenIssuerCmd).Standalone()

	ssoAdmin_updateTrustedTokenIssuerCmd.Flags().String("name", "", "Specifies the updated name to be applied to the trusted token issuer configuration.")
	ssoAdmin_updateTrustedTokenIssuerCmd.Flags().String("trusted-token-issuer-arn", "", "Specifies the ARN of the trusted token issuer configuration that you want to update.")
	ssoAdmin_updateTrustedTokenIssuerCmd.Flags().String("trusted-token-issuer-configuration", "", "Specifies a structure with settings to apply to the specified trusted token issuer.")
	ssoAdmin_updateTrustedTokenIssuerCmd.MarkFlagRequired("trusted-token-issuer-arn")
	ssoAdminCmd.AddCommand(ssoAdmin_updateTrustedTokenIssuerCmd)
}
