package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentcore_getWorkloadAccessTokenCmd = &cobra.Command{
	Use:   "get-workload-access-token",
	Short: "Obtains a workload access token for agentic workloads not acting on behalf of a user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentcore_getWorkloadAccessTokenCmd).Standalone()

	bedrockAgentcore_getWorkloadAccessTokenCmd.Flags().String("workload-name", "", "The unique identifier for the registered workload.")
	bedrockAgentcore_getWorkloadAccessTokenCmd.MarkFlagRequired("workload-name")
	bedrockAgentcoreCmd.AddCommand(bedrockAgentcore_getWorkloadAccessTokenCmd)
}
