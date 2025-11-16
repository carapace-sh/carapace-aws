package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentcoreControl_listApiKeyCredentialProvidersCmd = &cobra.Command{
	Use:   "list-api-key-credential-providers",
	Short: "Lists all API key credential providers in your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentcoreControl_listApiKeyCredentialProvidersCmd).Standalone()

	bedrockAgentcoreControl_listApiKeyCredentialProvidersCmd.Flags().String("max-results", "", "Maximum number of results to return.")
	bedrockAgentcoreControl_listApiKeyCredentialProvidersCmd.Flags().String("next-token", "", "Pagination token.")
	bedrockAgentcoreControlCmd.AddCommand(bedrockAgentcoreControl_listApiKeyCredentialProvidersCmd)
}
