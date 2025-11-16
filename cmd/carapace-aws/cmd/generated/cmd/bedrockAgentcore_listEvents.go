package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentcore_listEventsCmd = &cobra.Command{
	Use:   "list-events",
	Short: "Lists events in an AgentCore Memory resource based on specified criteria.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentcore_listEventsCmd).Standalone()

	bedrockAgentcore_listEventsCmd.Flags().String("actor-id", "", "The identifier of the actor for which to list events.")
	bedrockAgentcore_listEventsCmd.Flags().String("filter", "", "Filter criteria to apply when listing events.")
	bedrockAgentcore_listEventsCmd.Flags().Bool("include-payloads", false, "Specifies whether to include event payloads in the response.")
	bedrockAgentcore_listEventsCmd.Flags().String("max-results", "", "The maximum number of results to return in a single call.")
	bedrockAgentcore_listEventsCmd.Flags().String("memory-id", "", "The identifier of the AgentCore Memory resource for which to list events.")
	bedrockAgentcore_listEventsCmd.Flags().String("next-token", "", "The token for the next set of results.")
	bedrockAgentcore_listEventsCmd.Flags().Bool("no-include-payloads", false, "Specifies whether to include event payloads in the response.")
	bedrockAgentcore_listEventsCmd.Flags().String("session-id", "", "The identifier of the session for which to list events.")
	bedrockAgentcore_listEventsCmd.MarkFlagRequired("actor-id")
	bedrockAgentcore_listEventsCmd.MarkFlagRequired("memory-id")
	bedrockAgentcore_listEventsCmd.Flag("no-include-payloads").Hidden = true
	bedrockAgentcore_listEventsCmd.MarkFlagRequired("session-id")
	bedrockAgentcoreCmd.AddCommand(bedrockAgentcore_listEventsCmd)
}
