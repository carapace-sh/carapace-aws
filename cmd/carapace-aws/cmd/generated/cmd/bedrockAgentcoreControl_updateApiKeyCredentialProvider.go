package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentcoreControl_updateApiKeyCredentialProviderCmd = &cobra.Command{
	Use:   "update-api-key-credential-provider",
	Short: "Updates an existing API key credential provider.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentcoreControl_updateApiKeyCredentialProviderCmd).Standalone()

	bedrockAgentcoreControl_updateApiKeyCredentialProviderCmd.Flags().String("api-key", "", "The new API key to use for authentication.")
	bedrockAgentcoreControl_updateApiKeyCredentialProviderCmd.Flags().String("name", "", "The name of the API key credential provider to update.")
	bedrockAgentcoreControl_updateApiKeyCredentialProviderCmd.MarkFlagRequired("api-key")
	bedrockAgentcoreControl_updateApiKeyCredentialProviderCmd.MarkFlagRequired("name")
	bedrockAgentcoreControlCmd.AddCommand(bedrockAgentcoreControl_updateApiKeyCredentialProviderCmd)
}
