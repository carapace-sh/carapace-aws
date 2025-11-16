package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentRuntime_startFlowExecutionCmd = &cobra.Command{
	Use:   "start-flow-execution",
	Short: "Starts an execution of an Amazon Bedrock flow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentRuntime_startFlowExecutionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgentRuntime_startFlowExecutionCmd).Standalone()

		bedrockAgentRuntime_startFlowExecutionCmd.Flags().String("flow-alias-identifier", "", "The unique identifier of the flow alias to use for the flow execution.")
		bedrockAgentRuntime_startFlowExecutionCmd.Flags().String("flow-execution-name", "", "The unique name for the flow execution.")
		bedrockAgentRuntime_startFlowExecutionCmd.Flags().String("flow-identifier", "", "The unique identifier of the flow to execute.")
		bedrockAgentRuntime_startFlowExecutionCmd.Flags().String("inputs", "", "The input data required for the flow execution.")
		bedrockAgentRuntime_startFlowExecutionCmd.Flags().String("model-performance-configuration", "", "The performance settings for the foundation model used in the flow execution.")
		bedrockAgentRuntime_startFlowExecutionCmd.MarkFlagRequired("flow-alias-identifier")
		bedrockAgentRuntime_startFlowExecutionCmd.MarkFlagRequired("flow-identifier")
		bedrockAgentRuntime_startFlowExecutionCmd.MarkFlagRequired("inputs")
	})
	bedrockAgentRuntimeCmd.AddCommand(bedrockAgentRuntime_startFlowExecutionCmd)
}
