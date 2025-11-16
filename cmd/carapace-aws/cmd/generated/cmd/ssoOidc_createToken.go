package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssoOidc_createTokenCmd = &cobra.Command{
	Use:   "create-token",
	Short: "Creates and returns access and refresh tokens for clients that are authenticated using client secrets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssoOidc_createTokenCmd).Standalone()

	ssoOidc_createTokenCmd.Flags().String("client-id", "", "The unique identifier string for the client or application.")
	ssoOidc_createTokenCmd.Flags().String("client-secret", "", "A secret string generated for the client.")
	ssoOidc_createTokenCmd.Flags().String("code", "", "Used only when calling this API for the Authorization Code grant type.")
	ssoOidc_createTokenCmd.Flags().String("code-verifier", "", "Used only when calling this API for the Authorization Code grant type.")
	ssoOidc_createTokenCmd.Flags().String("device-code", "", "Used only when calling this API for the Device Code grant type.")
	ssoOidc_createTokenCmd.Flags().String("grant-type", "", "Supports the following OAuth grant types: Authorization Code, Device Code, and Refresh Token.")
	ssoOidc_createTokenCmd.Flags().String("redirect-uri", "", "Used only when calling this API for the Authorization Code grant type.")
	ssoOidc_createTokenCmd.Flags().String("refresh-token", "", "Used only when calling this API for the Refresh Token grant type.")
	ssoOidc_createTokenCmd.Flags().String("scope", "", "The list of scopes for which authorization is requested.")
	ssoOidc_createTokenCmd.MarkFlagRequired("client-id")
	ssoOidc_createTokenCmd.MarkFlagRequired("client-secret")
	ssoOidc_createTokenCmd.MarkFlagRequired("grant-type")
	ssoOidcCmd.AddCommand(ssoOidc_createTokenCmd)
}
