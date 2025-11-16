package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentcore_completeResourceTokenAuthCmd = &cobra.Command{
	Use:   "complete-resource-token-auth",
	Short: "Confirms the user authentication session for obtaining OAuth2.0 tokens for a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentcore_completeResourceTokenAuthCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgentcore_completeResourceTokenAuthCmd).Standalone()

		bedrockAgentcore_completeResourceTokenAuthCmd.Flags().String("session-uri", "", "Unique identifier for the user's authentication session for retrieving OAuth2 tokens.")
		bedrockAgentcore_completeResourceTokenAuthCmd.Flags().String("user-identifier", "", "The OAuth2.0 token or user ID that was used to generate the workload access token used for initiating the user authorization flow to retrieve OAuth2.0 tokens.")
		bedrockAgentcore_completeResourceTokenAuthCmd.MarkFlagRequired("session-uri")
		bedrockAgentcore_completeResourceTokenAuthCmd.MarkFlagRequired("user-identifier")
	})
	bedrockAgentcoreCmd.AddCommand(bedrockAgentcore_completeResourceTokenAuthCmd)
}
