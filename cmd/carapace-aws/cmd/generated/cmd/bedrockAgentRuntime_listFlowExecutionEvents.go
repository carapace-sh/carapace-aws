package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentRuntime_listFlowExecutionEventsCmd = &cobra.Command{
	Use:   "list-flow-execution-events",
	Short: "Lists events that occurred during a flow execution.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentRuntime_listFlowExecutionEventsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgentRuntime_listFlowExecutionEventsCmd).Standalone()

		bedrockAgentRuntime_listFlowExecutionEventsCmd.Flags().String("event-type", "", "The type of events to retrieve.")
		bedrockAgentRuntime_listFlowExecutionEventsCmd.Flags().String("execution-identifier", "", "The unique identifier of the flow execution.")
		bedrockAgentRuntime_listFlowExecutionEventsCmd.Flags().String("flow-alias-identifier", "", "The unique identifier of the flow alias used for the execution.")
		bedrockAgentRuntime_listFlowExecutionEventsCmd.Flags().String("flow-identifier", "", "The unique identifier of the flow.")
		bedrockAgentRuntime_listFlowExecutionEventsCmd.Flags().String("max-results", "", "The maximum number of events to return in a single response.")
		bedrockAgentRuntime_listFlowExecutionEventsCmd.Flags().String("next-token", "", "A token to retrieve the next set of results.")
		bedrockAgentRuntime_listFlowExecutionEventsCmd.MarkFlagRequired("event-type")
		bedrockAgentRuntime_listFlowExecutionEventsCmd.MarkFlagRequired("execution-identifier")
		bedrockAgentRuntime_listFlowExecutionEventsCmd.MarkFlagRequired("flow-alias-identifier")
		bedrockAgentRuntime_listFlowExecutionEventsCmd.MarkFlagRequired("flow-identifier")
	})
	bedrockAgentRuntimeCmd.AddCommand(bedrockAgentRuntime_listFlowExecutionEventsCmd)
}
