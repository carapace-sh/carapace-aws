package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentcore_createEventCmd = &cobra.Command{
	Use:   "create-event",
	Short: "Creates an event in an AgentCore Memory resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentcore_createEventCmd).Standalone()

	bedrockAgentcore_createEventCmd.Flags().String("actor-id", "", "The identifier of the actor associated with this event.")
	bedrockAgentcore_createEventCmd.Flags().String("branch", "", "The branch information for this event.")
	bedrockAgentcore_createEventCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier to ensure that the operation completes no more than one time.")
	bedrockAgentcore_createEventCmd.Flags().String("event-timestamp", "", "The timestamp when the event occurred.")
	bedrockAgentcore_createEventCmd.Flags().String("memory-id", "", "The identifier of the AgentCore Memory resource in which to create the event.")
	bedrockAgentcore_createEventCmd.Flags().String("metadata", "", "The key-value metadata to attach to the event.")
	bedrockAgentcore_createEventCmd.Flags().String("payload", "", "The content payload of the event.")
	bedrockAgentcore_createEventCmd.Flags().String("session-id", "", "The identifier of the session in which this event occurs.")
	bedrockAgentcore_createEventCmd.MarkFlagRequired("actor-id")
	bedrockAgentcore_createEventCmd.MarkFlagRequired("event-timestamp")
	bedrockAgentcore_createEventCmd.MarkFlagRequired("memory-id")
	bedrockAgentcore_createEventCmd.MarkFlagRequired("payload")
	bedrockAgentcoreCmd.AddCommand(bedrockAgentcore_createEventCmd)
}
