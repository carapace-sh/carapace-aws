package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentcore_getResourceOauth2TokenCmd = &cobra.Command{
	Use:   "get-resource-oauth2-token",
	Short: "Returns the OAuth 2.0 token of the provided resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentcore_getResourceOauth2TokenCmd).Standalone()

	bedrockAgentcore_getResourceOauth2TokenCmd.Flags().String("custom-parameters", "", "A map of custom parameters to include in the authorization request to the resource credential provider.")
	bedrockAgentcore_getResourceOauth2TokenCmd.Flags().String("custom-state", "", "An opaque string that will be sent back to the callback URL provided in resourceOauth2ReturnUrl.")
	bedrockAgentcore_getResourceOauth2TokenCmd.Flags().Bool("force-authentication", false, "Indicates whether to always initiate a new three-legged OAuth (3LO) flow, regardless of any existing session.")
	bedrockAgentcore_getResourceOauth2TokenCmd.Flags().Bool("no-force-authentication", false, "Indicates whether to always initiate a new three-legged OAuth (3LO) flow, regardless of any existing session.")
	bedrockAgentcore_getResourceOauth2TokenCmd.Flags().String("oauth2-flow", "", "The type of flow to be performed.")
	bedrockAgentcore_getResourceOauth2TokenCmd.Flags().String("resource-credential-provider-name", "", "The name of the resource's credential provider.")
	bedrockAgentcore_getResourceOauth2TokenCmd.Flags().String("resource-oauth2-return-url", "", "The callback URL to redirect to after the OAuth 2.0 token retrieval is complete.")
	bedrockAgentcore_getResourceOauth2TokenCmd.Flags().String("scopes", "", "The OAuth scopes being requested.")
	bedrockAgentcore_getResourceOauth2TokenCmd.Flags().String("session-uri", "", "Unique identifier for the user's authentication session for retrieving OAuth2 tokens.")
	bedrockAgentcore_getResourceOauth2TokenCmd.Flags().String("workload-identity-token", "", "The identity token of the workload from which you want to retrieve the OAuth2 token.")
	bedrockAgentcore_getResourceOauth2TokenCmd.Flag("no-force-authentication").Hidden = true
	bedrockAgentcore_getResourceOauth2TokenCmd.MarkFlagRequired("oauth2-flow")
	bedrockAgentcore_getResourceOauth2TokenCmd.MarkFlagRequired("resource-credential-provider-name")
	bedrockAgentcore_getResourceOauth2TokenCmd.MarkFlagRequired("scopes")
	bedrockAgentcore_getResourceOauth2TokenCmd.MarkFlagRequired("workload-identity-token")
	bedrockAgentcoreCmd.AddCommand(bedrockAgentcore_getResourceOauth2TokenCmd)
}
