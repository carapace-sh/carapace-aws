package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentcore_getWorkloadAccessTokenForJwtCmd = &cobra.Command{
	Use:   "get-workload-access-token-for-jwt",
	Short: "Obtains a workload access token for agentic workloads acting on behalf of a user, using a JWT token.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentcore_getWorkloadAccessTokenForJwtCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgentcore_getWorkloadAccessTokenForJwtCmd).Standalone()

		bedrockAgentcore_getWorkloadAccessTokenForJwtCmd.Flags().String("user-token", "", "The OAuth 2.0 token issued by the user's identity provider.")
		bedrockAgentcore_getWorkloadAccessTokenForJwtCmd.Flags().String("workload-name", "", "The unique identifier for the registered workload.")
		bedrockAgentcore_getWorkloadAccessTokenForJwtCmd.MarkFlagRequired("user-token")
		bedrockAgentcore_getWorkloadAccessTokenForJwtCmd.MarkFlagRequired("workload-name")
	})
	bedrockAgentcoreCmd.AddCommand(bedrockAgentcore_getWorkloadAccessTokenForJwtCmd)
}
