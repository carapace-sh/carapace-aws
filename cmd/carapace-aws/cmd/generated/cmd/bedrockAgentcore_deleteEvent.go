package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentcore_deleteEventCmd = &cobra.Command{
	Use:   "delete-event",
	Short: "Deletes an event from an AgentCore Memory resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentcore_deleteEventCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgentcore_deleteEventCmd).Standalone()

		bedrockAgentcore_deleteEventCmd.Flags().String("actor-id", "", "The identifier of the actor associated with the event to delete.")
		bedrockAgentcore_deleteEventCmd.Flags().String("event-id", "", "The identifier of the event to delete.")
		bedrockAgentcore_deleteEventCmd.Flags().String("memory-id", "", "The identifier of the AgentCore Memory resource from which to delete the event.")
		bedrockAgentcore_deleteEventCmd.Flags().String("session-id", "", "The identifier of the session containing the event to delete.")
		bedrockAgentcore_deleteEventCmd.MarkFlagRequired("actor-id")
		bedrockAgentcore_deleteEventCmd.MarkFlagRequired("event-id")
		bedrockAgentcore_deleteEventCmd.MarkFlagRequired("memory-id")
		bedrockAgentcore_deleteEventCmd.MarkFlagRequired("session-id")
	})
	bedrockAgentcoreCmd.AddCommand(bedrockAgentcore_deleteEventCmd)
}
