package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgent_listAgentAliasesCmd = &cobra.Command{
	Use:   "list-agent-aliases",
	Short: "Lists the aliases of an agent and information about each one.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgent_listAgentAliasesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgent_listAgentAliasesCmd).Standalone()

		bedrockAgent_listAgentAliasesCmd.Flags().String("agent-id", "", "The unique identifier of the agent.")
		bedrockAgent_listAgentAliasesCmd.Flags().String("max-results", "", "The maximum number of results to return in the response.")
		bedrockAgent_listAgentAliasesCmd.Flags().String("next-token", "", "If the total number of results is greater than the `maxResults` value provided in the request, enter the token returned in the `nextToken` field in the response in this field to return the next batch of results.")
		bedrockAgent_listAgentAliasesCmd.MarkFlagRequired("agent-id")
	})
	bedrockAgentCmd.AddCommand(bedrockAgent_listAgentAliasesCmd)
}
