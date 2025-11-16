package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssoAdmin_createTrustedTokenIssuerCmd = &cobra.Command{
	Use:   "create-trusted-token-issuer",
	Short: "Creates a connection to a trusted token issuer in an instance of IAM Identity Center.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssoAdmin_createTrustedTokenIssuerCmd).Standalone()

	ssoAdmin_createTrustedTokenIssuerCmd.Flags().String("client-token", "", "Specifies a unique, case-sensitive ID that you provide to ensure the idempotency of the request.")
	ssoAdmin_createTrustedTokenIssuerCmd.Flags().String("instance-arn", "", "Specifies the ARN of the instance of IAM Identity Center to contain the new trusted token issuer configuration.")
	ssoAdmin_createTrustedTokenIssuerCmd.Flags().String("name", "", "Specifies the name of the new trusted token issuer configuration.")
	ssoAdmin_createTrustedTokenIssuerCmd.Flags().String("tags", "", "Specifies tags to be attached to the new trusted token issuer configuration.")
	ssoAdmin_createTrustedTokenIssuerCmd.Flags().String("trusted-token-issuer-configuration", "", "Specifies settings that apply to the new trusted token issuer configuration.")
	ssoAdmin_createTrustedTokenIssuerCmd.Flags().String("trusted-token-issuer-type", "", "Specifies the type of the new trusted token issuer.")
	ssoAdmin_createTrustedTokenIssuerCmd.MarkFlagRequired("instance-arn")
	ssoAdmin_createTrustedTokenIssuerCmd.MarkFlagRequired("name")
	ssoAdmin_createTrustedTokenIssuerCmd.MarkFlagRequired("trusted-token-issuer-configuration")
	ssoAdmin_createTrustedTokenIssuerCmd.MarkFlagRequired("trusted-token-issuer-type")
	ssoAdminCmd.AddCommand(ssoAdmin_createTrustedTokenIssuerCmd)
}
