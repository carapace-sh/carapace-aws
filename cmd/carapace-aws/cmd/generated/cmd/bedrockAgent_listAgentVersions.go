package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgent_listAgentVersionsCmd = &cobra.Command{
	Use:   "list-agent-versions",
	Short: "Lists the versions of an agent and information about each version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgent_listAgentVersionsCmd).Standalone()

	bedrockAgent_listAgentVersionsCmd.Flags().String("agent-id", "", "The unique identifier of the agent.")
	bedrockAgent_listAgentVersionsCmd.Flags().String("max-results", "", "The maximum number of results to return in the response.")
	bedrockAgent_listAgentVersionsCmd.Flags().String("next-token", "", "If the total number of results is greater than the `maxResults` value provided in the request, enter the token returned in the `nextToken` field in the response in this field to return the next batch of results.")
	bedrockAgent_listAgentVersionsCmd.MarkFlagRequired("agent-id")
	bedrockAgentCmd.AddCommand(bedrockAgent_listAgentVersionsCmd)
}
