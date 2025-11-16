package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgent_listAgentActionGroupsCmd = &cobra.Command{
	Use:   "list-agent-action-groups",
	Short: "Lists the action groups for an agent and information about each one.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgent_listAgentActionGroupsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgent_listAgentActionGroupsCmd).Standalone()

		bedrockAgent_listAgentActionGroupsCmd.Flags().String("agent-id", "", "The unique identifier of the agent.")
		bedrockAgent_listAgentActionGroupsCmd.Flags().String("agent-version", "", "The version of the agent.")
		bedrockAgent_listAgentActionGroupsCmd.Flags().String("max-results", "", "The maximum number of results to return in the response.")
		bedrockAgent_listAgentActionGroupsCmd.Flags().String("next-token", "", "If the total number of results is greater than the `maxResults` value provided in the request, enter the token returned in the `nextToken` field in the response in this field to return the next batch of results.")
		bedrockAgent_listAgentActionGroupsCmd.MarkFlagRequired("agent-id")
		bedrockAgent_listAgentActionGroupsCmd.MarkFlagRequired("agent-version")
	})
	bedrockAgentCmd.AddCommand(bedrockAgent_listAgentActionGroupsCmd)
}
