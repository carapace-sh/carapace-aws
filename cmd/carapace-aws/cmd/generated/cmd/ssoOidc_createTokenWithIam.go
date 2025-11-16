package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssoOidc_createTokenWithIamCmd = &cobra.Command{
	Use:   "create-token-with-iam",
	Short: "Creates and returns access and refresh tokens for authorized client applications that are authenticated using any IAM entity, such as a service role or user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssoOidc_createTokenWithIamCmd).Standalone()

	ssoOidc_createTokenWithIamCmd.Flags().String("assertion", "", "Used only when calling this API for the JWT Bearer grant type.")
	ssoOidc_createTokenWithIamCmd.Flags().String("client-id", "", "The unique identifier string for the client or application.")
	ssoOidc_createTokenWithIamCmd.Flags().String("code", "", "Used only when calling this API for the Authorization Code grant type.")
	ssoOidc_createTokenWithIamCmd.Flags().String("code-verifier", "", "Used only when calling this API for the Authorization Code grant type.")
	ssoOidc_createTokenWithIamCmd.Flags().String("grant-type", "", "Supports the following OAuth grant types: Authorization Code, Refresh Token, JWT Bearer, and Token Exchange.")
	ssoOidc_createTokenWithIamCmd.Flags().String("redirect-uri", "", "Used only when calling this API for the Authorization Code grant type.")
	ssoOidc_createTokenWithIamCmd.Flags().String("refresh-token", "", "Used only when calling this API for the Refresh Token grant type.")
	ssoOidc_createTokenWithIamCmd.Flags().String("requested-token-type", "", "Used only when calling this API for the Token Exchange grant type.")
	ssoOidc_createTokenWithIamCmd.Flags().String("scope", "", "The list of scopes for which authorization is requested.")
	ssoOidc_createTokenWithIamCmd.Flags().String("subject-token", "", "Used only when calling this API for the Token Exchange grant type.")
	ssoOidc_createTokenWithIamCmd.Flags().String("subject-token-type", "", "Used only when calling this API for the Token Exchange grant type.")
	ssoOidc_createTokenWithIamCmd.MarkFlagRequired("client-id")
	ssoOidc_createTokenWithIamCmd.MarkFlagRequired("grant-type")
	ssoOidcCmd.AddCommand(ssoOidc_createTokenWithIamCmd)
}
