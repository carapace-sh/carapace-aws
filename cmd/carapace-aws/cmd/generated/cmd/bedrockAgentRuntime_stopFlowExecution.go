package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentRuntime_stopFlowExecutionCmd = &cobra.Command{
	Use:   "stop-flow-execution",
	Short: "Stops an Amazon Bedrock flow's execution.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentRuntime_stopFlowExecutionCmd).Standalone()

	bedrockAgentRuntime_stopFlowExecutionCmd.Flags().String("execution-identifier", "", "The unique identifier of the flow execution to stop.")
	bedrockAgentRuntime_stopFlowExecutionCmd.Flags().String("flow-alias-identifier", "", "The unique identifier of the flow alias used for the execution.")
	bedrockAgentRuntime_stopFlowExecutionCmd.Flags().String("flow-identifier", "", "The unique identifier of the flow.")
	bedrockAgentRuntime_stopFlowExecutionCmd.MarkFlagRequired("execution-identifier")
	bedrockAgentRuntime_stopFlowExecutionCmd.MarkFlagRequired("flow-alias-identifier")
	bedrockAgentRuntime_stopFlowExecutionCmd.MarkFlagRequired("flow-identifier")
	bedrockAgentRuntimeCmd.AddCommand(bedrockAgentRuntime_stopFlowExecutionCmd)
}
