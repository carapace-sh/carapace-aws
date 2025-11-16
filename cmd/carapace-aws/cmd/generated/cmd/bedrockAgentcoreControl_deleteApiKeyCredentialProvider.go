package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentcoreControl_deleteApiKeyCredentialProviderCmd = &cobra.Command{
	Use:   "delete-api-key-credential-provider",
	Short: "Deletes an API key credential provider.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentcoreControl_deleteApiKeyCredentialProviderCmd).Standalone()

	bedrockAgentcoreControl_deleteApiKeyCredentialProviderCmd.Flags().String("name", "", "The name of the API key credential provider to delete.")
	bedrockAgentcoreControl_deleteApiKeyCredentialProviderCmd.MarkFlagRequired("name")
	bedrockAgentcoreControlCmd.AddCommand(bedrockAgentcoreControl_deleteApiKeyCredentialProviderCmd)
}
