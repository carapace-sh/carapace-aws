package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentcoreControl_getOauth2CredentialProviderCmd = &cobra.Command{
	Use:   "get-oauth2-credential-provider",
	Short: "Retrieves information about an OAuth2 credential provider.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentcoreControl_getOauth2CredentialProviderCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgentcoreControl_getOauth2CredentialProviderCmd).Standalone()

		bedrockAgentcoreControl_getOauth2CredentialProviderCmd.Flags().String("name", "", "The name of the OAuth2 credential provider to retrieve.")
		bedrockAgentcoreControl_getOauth2CredentialProviderCmd.MarkFlagRequired("name")
	})
	bedrockAgentcoreControlCmd.AddCommand(bedrockAgentcoreControl_getOauth2CredentialProviderCmd)
}
