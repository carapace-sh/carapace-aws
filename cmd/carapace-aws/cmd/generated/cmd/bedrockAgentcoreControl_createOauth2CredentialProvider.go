package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentcoreControl_createOauth2CredentialProviderCmd = &cobra.Command{
	Use:   "create-oauth2-credential-provider",
	Short: "Creates a new OAuth2 credential provider.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentcoreControl_createOauth2CredentialProviderCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgentcoreControl_createOauth2CredentialProviderCmd).Standalone()

		bedrockAgentcoreControl_createOauth2CredentialProviderCmd.Flags().String("credential-provider-vendor", "", "The vendor of the OAuth2 credential provider.")
		bedrockAgentcoreControl_createOauth2CredentialProviderCmd.Flags().String("name", "", "The name of the OAuth2 credential provider.")
		bedrockAgentcoreControl_createOauth2CredentialProviderCmd.Flags().String("oauth2-provider-config-input", "", "The configuration settings for the OAuth2 provider, including client ID, client secret, and other vendor-specific settings.")
		bedrockAgentcoreControl_createOauth2CredentialProviderCmd.Flags().String("tags", "", "A map of tag keys and values to assign to the OAuth2 credential provider.")
		bedrockAgentcoreControl_createOauth2CredentialProviderCmd.MarkFlagRequired("credential-provider-vendor")
		bedrockAgentcoreControl_createOauth2CredentialProviderCmd.MarkFlagRequired("name")
		bedrockAgentcoreControl_createOauth2CredentialProviderCmd.MarkFlagRequired("oauth2-provider-config-input")
	})
	bedrockAgentcoreControlCmd.AddCommand(bedrockAgentcoreControl_createOauth2CredentialProviderCmd)
}
