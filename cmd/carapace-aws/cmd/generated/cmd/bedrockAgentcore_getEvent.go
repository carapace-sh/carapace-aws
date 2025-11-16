package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentcore_getEventCmd = &cobra.Command{
	Use:   "get-event",
	Short: "Retrieves information about a specific event in an AgentCore Memory resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentcore_getEventCmd).Standalone()

	bedrockAgentcore_getEventCmd.Flags().String("actor-id", "", "The identifier of the actor associated with the event.")
	bedrockAgentcore_getEventCmd.Flags().String("event-id", "", "The identifier of the event to retrieve.")
	bedrockAgentcore_getEventCmd.Flags().String("memory-id", "", "The identifier of the AgentCore Memory resource containing the event.")
	bedrockAgentcore_getEventCmd.Flags().String("session-id", "", "The identifier of the session containing the event.")
	bedrockAgentcore_getEventCmd.MarkFlagRequired("actor-id")
	bedrockAgentcore_getEventCmd.MarkFlagRequired("event-id")
	bedrockAgentcore_getEventCmd.MarkFlagRequired("memory-id")
	bedrockAgentcore_getEventCmd.MarkFlagRequired("session-id")
	bedrockAgentcoreCmd.AddCommand(bedrockAgentcore_getEventCmd)
}
