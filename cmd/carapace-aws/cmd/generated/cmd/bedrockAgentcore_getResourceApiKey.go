package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentcore_getResourceApiKeyCmd = &cobra.Command{
	Use:   "get-resource-api-key",
	Short: "Retrieves the API key associated with an API key credential provider.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentcore_getResourceApiKeyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgentcore_getResourceApiKeyCmd).Standalone()

		bedrockAgentcore_getResourceApiKeyCmd.Flags().String("resource-credential-provider-name", "", "The credential provider name for the resource from which you are retrieving the API key.")
		bedrockAgentcore_getResourceApiKeyCmd.Flags().String("workload-identity-token", "", "The identity token of the workload from which you want to retrieve the API key.")
		bedrockAgentcore_getResourceApiKeyCmd.MarkFlagRequired("resource-credential-provider-name")
		bedrockAgentcore_getResourceApiKeyCmd.MarkFlagRequired("workload-identity-token")
	})
	bedrockAgentcoreCmd.AddCommand(bedrockAgentcore_getResourceApiKeyCmd)
}
