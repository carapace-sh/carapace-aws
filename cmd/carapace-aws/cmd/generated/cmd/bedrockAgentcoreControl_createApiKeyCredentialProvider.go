package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentcoreControl_createApiKeyCredentialProviderCmd = &cobra.Command{
	Use:   "create-api-key-credential-provider",
	Short: "Creates a new API key credential provider.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentcoreControl_createApiKeyCredentialProviderCmd).Standalone()

	bedrockAgentcoreControl_createApiKeyCredentialProviderCmd.Flags().String("api-key", "", "The API key to use for authentication.")
	bedrockAgentcoreControl_createApiKeyCredentialProviderCmd.Flags().String("name", "", "The name of the API key credential provider.")
	bedrockAgentcoreControl_createApiKeyCredentialProviderCmd.Flags().String("tags", "", "A map of tag keys and values to assign to the API key credential provider.")
	bedrockAgentcoreControl_createApiKeyCredentialProviderCmd.MarkFlagRequired("api-key")
	bedrockAgentcoreControl_createApiKeyCredentialProviderCmd.MarkFlagRequired("name")
	bedrockAgentcoreControlCmd.AddCommand(bedrockAgentcoreControl_createApiKeyCredentialProviderCmd)
}
