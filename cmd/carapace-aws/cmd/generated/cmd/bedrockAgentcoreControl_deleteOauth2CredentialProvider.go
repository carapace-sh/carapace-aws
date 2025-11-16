package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentcoreControl_deleteOauth2CredentialProviderCmd = &cobra.Command{
	Use:   "delete-oauth2-credential-provider",
	Short: "Deletes an OAuth2 credential provider.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentcoreControl_deleteOauth2CredentialProviderCmd).Standalone()

	bedrockAgentcoreControl_deleteOauth2CredentialProviderCmd.Flags().String("name", "", "The name of the OAuth2 credential provider to delete.")
	bedrockAgentcoreControl_deleteOauth2CredentialProviderCmd.MarkFlagRequired("name")
	bedrockAgentcoreControlCmd.AddCommand(bedrockAgentcoreControl_deleteOauth2CredentialProviderCmd)
}
