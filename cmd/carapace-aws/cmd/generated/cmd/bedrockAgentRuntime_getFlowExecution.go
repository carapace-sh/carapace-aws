package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentRuntime_getFlowExecutionCmd = &cobra.Command{
	Use:   "get-flow-execution",
	Short: "Retrieves details about a specific flow execution, including its status, start and end times, and any errors that occurred during execution.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentRuntime_getFlowExecutionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgentRuntime_getFlowExecutionCmd).Standalone()

		bedrockAgentRuntime_getFlowExecutionCmd.Flags().String("execution-identifier", "", "The unique identifier of the flow execution to retrieve.")
		bedrockAgentRuntime_getFlowExecutionCmd.Flags().String("flow-alias-identifier", "", "The unique identifier of the flow alias used for the execution.")
		bedrockAgentRuntime_getFlowExecutionCmd.Flags().String("flow-identifier", "", "The unique identifier of the flow.")
		bedrockAgentRuntime_getFlowExecutionCmd.MarkFlagRequired("execution-identifier")
		bedrockAgentRuntime_getFlowExecutionCmd.MarkFlagRequired("flow-alias-identifier")
		bedrockAgentRuntime_getFlowExecutionCmd.MarkFlagRequired("flow-identifier")
	})
	bedrockAgentRuntimeCmd.AddCommand(bedrockAgentRuntime_getFlowExecutionCmd)
}
