package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssoOidc_registerClientCmd = &cobra.Command{
	Use:   "register-client",
	Short: "Registers a public\u00a0client with IAM Identity Center.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssoOidc_registerClientCmd).Standalone()

	ssoOidc_registerClientCmd.Flags().String("client-name", "", "The friendly name of the client.")
	ssoOidc_registerClientCmd.Flags().String("client-type", "", "The type of client.")
	ssoOidc_registerClientCmd.Flags().String("entitled-application-arn", "", "This IAM Identity Center application ARN is used to define administrator-managed configuration for public client access to resources.")
	ssoOidc_registerClientCmd.Flags().String("grant-types", "", "The list of OAuth 2.0 grant types that are defined by the client.")
	ssoOidc_registerClientCmd.Flags().String("issuer-url", "", "The IAM Identity Center Issuer URL associated with an instance of IAM Identity Center.")
	ssoOidc_registerClientCmd.Flags().String("redirect-uris", "", "The list of redirect URI that are defined by the client.")
	ssoOidc_registerClientCmd.Flags().String("scopes", "", "The list of scopes that are defined by the client.")
	ssoOidc_registerClientCmd.MarkFlagRequired("client-name")
	ssoOidc_registerClientCmd.MarkFlagRequired("client-type")
	ssoOidcCmd.AddCommand(ssoOidc_registerClientCmd)
}
