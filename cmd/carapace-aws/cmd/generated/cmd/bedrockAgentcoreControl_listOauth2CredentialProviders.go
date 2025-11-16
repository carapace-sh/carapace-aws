package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentcoreControl_listOauth2CredentialProvidersCmd = &cobra.Command{
	Use:   "list-oauth2-credential-providers",
	Short: "Lists all OAuth2 credential providers in your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentcoreControl_listOauth2CredentialProvidersCmd).Standalone()

	bedrockAgentcoreControl_listOauth2CredentialProvidersCmd.Flags().String("max-results", "", "Maximum number of results to return.")
	bedrockAgentcoreControl_listOauth2CredentialProvidersCmd.Flags().String("next-token", "", "Pagination token.")
	bedrockAgentcoreControlCmd.AddCommand(bedrockAgentcoreControl_listOauth2CredentialProvidersCmd)
}
