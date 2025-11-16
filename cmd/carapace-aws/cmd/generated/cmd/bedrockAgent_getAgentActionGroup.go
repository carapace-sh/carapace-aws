package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgent_getAgentActionGroupCmd = &cobra.Command{
	Use:   "get-agent-action-group",
	Short: "Gets information about an action group for an agent.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgent_getAgentActionGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgent_getAgentActionGroupCmd).Standalone()

		bedrockAgent_getAgentActionGroupCmd.Flags().String("action-group-id", "", "The unique identifier of the action group for which to get information.")
		bedrockAgent_getAgentActionGroupCmd.Flags().String("agent-id", "", "The unique identifier of the agent that the action group belongs to.")
		bedrockAgent_getAgentActionGroupCmd.Flags().String("agent-version", "", "The version of the agent that the action group belongs to.")
		bedrockAgent_getAgentActionGroupCmd.MarkFlagRequired("action-group-id")
		bedrockAgent_getAgentActionGroupCmd.MarkFlagRequired("agent-id")
		bedrockAgent_getAgentActionGroupCmd.MarkFlagRequired("agent-version")
	})
	bedrockAgentCmd.AddCommand(bedrockAgent_getAgentActionGroupCmd)
}
