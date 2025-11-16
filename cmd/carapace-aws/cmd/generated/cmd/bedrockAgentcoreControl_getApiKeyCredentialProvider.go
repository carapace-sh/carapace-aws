package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentcoreControl_getApiKeyCredentialProviderCmd = &cobra.Command{
	Use:   "get-api-key-credential-provider",
	Short: "Retrieves information about an API key credential provider.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentcoreControl_getApiKeyCredentialProviderCmd).Standalone()

	bedrockAgentcoreControl_getApiKeyCredentialProviderCmd.Flags().String("name", "", "The name of the API key credential provider to retrieve.")
	bedrockAgentcoreControl_getApiKeyCredentialProviderCmd.MarkFlagRequired("name")
	bedrockAgentcoreControlCmd.AddCommand(bedrockAgentcoreControl_getApiKeyCredentialProviderCmd)
}
