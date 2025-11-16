package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentcore_getWorkloadAccessTokenForUserIdCmd = &cobra.Command{
	Use:   "get-workload-access-token-for-user-id",
	Short: "Obtains a workload access token for agentic workloads acting on behalf of a user, using the user's ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentcore_getWorkloadAccessTokenForUserIdCmd).Standalone()

	bedrockAgentcore_getWorkloadAccessTokenForUserIdCmd.Flags().String("user-id", "", "The ID of the user for whom you are retrieving the access token.")
	bedrockAgentcore_getWorkloadAccessTokenForUserIdCmd.Flags().String("workload-name", "", "The name of the workload from which you want to retrieve the access token.")
	bedrockAgentcore_getWorkloadAccessTokenForUserIdCmd.MarkFlagRequired("user-id")
	bedrockAgentcore_getWorkloadAccessTokenForUserIdCmd.MarkFlagRequired("workload-name")
	bedrockAgentcoreCmd.AddCommand(bedrockAgentcore_getWorkloadAccessTokenForUserIdCmd)
}
