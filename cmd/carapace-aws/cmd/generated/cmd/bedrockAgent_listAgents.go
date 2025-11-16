package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgent_listAgentsCmd = &cobra.Command{
	Use:   "list-agents",
	Short: "Lists the agents belonging to an account and information about each agent.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgent_listAgentsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgent_listAgentsCmd).Standalone()

		bedrockAgent_listAgentsCmd.Flags().String("max-results", "", "The maximum number of results to return in the response.")
		bedrockAgent_listAgentsCmd.Flags().String("next-token", "", "If the total number of results is greater than the `maxResults` value provided in the request, enter the token returned in the `nextToken` field in the response in this field to return the next batch of results.")
	})
	bedrockAgentCmd.AddCommand(bedrockAgent_listAgentsCmd)
}
