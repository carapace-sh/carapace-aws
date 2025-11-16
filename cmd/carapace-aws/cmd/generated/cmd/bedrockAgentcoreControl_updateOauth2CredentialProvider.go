package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentcoreControl_updateOauth2CredentialProviderCmd = &cobra.Command{
	Use:   "update-oauth2-credential-provider",
	Short: "Updates an existing OAuth2 credential provider.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentcoreControl_updateOauth2CredentialProviderCmd).Standalone()

	bedrockAgentcoreControl_updateOauth2CredentialProviderCmd.Flags().String("credential-provider-vendor", "", "The vendor of the OAuth2 credential provider.")
	bedrockAgentcoreControl_updateOauth2CredentialProviderCmd.Flags().String("name", "", "The name of the OAuth2 credential provider to update.")
	bedrockAgentcoreControl_updateOauth2CredentialProviderCmd.Flags().String("oauth2-provider-config-input", "", "The configuration input for the OAuth2 provider.")
	bedrockAgentcoreControl_updateOauth2CredentialProviderCmd.MarkFlagRequired("credential-provider-vendor")
	bedrockAgentcoreControl_updateOauth2CredentialProviderCmd.MarkFlagRequired("name")
	bedrockAgentcoreControl_updateOauth2CredentialProviderCmd.MarkFlagRequired("oauth2-provider-config-input")
	bedrockAgentcoreControlCmd.AddCommand(bedrockAgentcoreControl_updateOauth2CredentialProviderCmd)
}
