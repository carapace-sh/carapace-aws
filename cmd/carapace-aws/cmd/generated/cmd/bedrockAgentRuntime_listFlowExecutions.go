package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentRuntime_listFlowExecutionsCmd = &cobra.Command{
	Use:   "list-flow-executions",
	Short: "Lists all executions of a flow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentRuntime_listFlowExecutionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgentRuntime_listFlowExecutionsCmd).Standalone()

		bedrockAgentRuntime_listFlowExecutionsCmd.Flags().String("flow-alias-identifier", "", "The unique identifier of the flow alias to list executions for.")
		bedrockAgentRuntime_listFlowExecutionsCmd.Flags().String("flow-identifier", "", "The unique identifier of the flow to list executions for.")
		bedrockAgentRuntime_listFlowExecutionsCmd.Flags().String("max-results", "", "The maximum number of flow executions to return in a single response.")
		bedrockAgentRuntime_listFlowExecutionsCmd.Flags().String("next-token", "", "A token to retrieve the next set of results.")
		bedrockAgentRuntime_listFlowExecutionsCmd.MarkFlagRequired("flow-identifier")
	})
	bedrockAgentRuntimeCmd.AddCommand(bedrockAgentRuntime_listFlowExecutionsCmd)
}
