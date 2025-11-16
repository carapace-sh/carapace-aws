package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentcore_listActorsCmd = &cobra.Command{
	Use:   "list-actors",
	Short: "Lists all actors in an AgentCore Memory resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentcore_listActorsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgentcore_listActorsCmd).Standalone()

		bedrockAgentcore_listActorsCmd.Flags().String("max-results", "", "The maximum number of results to return in a single call.")
		bedrockAgentcore_listActorsCmd.Flags().String("memory-id", "", "The identifier of the AgentCore Memory resource for which to list actors.")
		bedrockAgentcore_listActorsCmd.Flags().String("next-token", "", "The token for the next set of results.")
		bedrockAgentcore_listActorsCmd.MarkFlagRequired("memory-id")
	})
	bedrockAgentcoreCmd.AddCommand(bedrockAgentcore_listActorsCmd)
}
