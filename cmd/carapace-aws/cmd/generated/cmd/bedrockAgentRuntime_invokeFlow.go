package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentRuntime_invokeFlowCmd = &cobra.Command{
	Use:   "invoke-flow",
	Short: "Invokes an alias of a flow to run the inputs that you specify and return the output of each node as a stream.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentRuntime_invokeFlowCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgentRuntime_invokeFlowCmd).Standalone()

		bedrockAgentRuntime_invokeFlowCmd.Flags().Bool("enable-trace", false, "Specifies whether to return the trace for the flow or not.")
		bedrockAgentRuntime_invokeFlowCmd.Flags().String("execution-id", "", "The unique identifier for the current flow execution.")
		bedrockAgentRuntime_invokeFlowCmd.Flags().String("flow-alias-identifier", "", "The unique identifier of the flow alias.")
		bedrockAgentRuntime_invokeFlowCmd.Flags().String("flow-identifier", "", "The unique identifier of the flow.")
		bedrockAgentRuntime_invokeFlowCmd.Flags().String("inputs", "", "A list of objects, each containing information about an input into the flow.")
		bedrockAgentRuntime_invokeFlowCmd.Flags().String("model-performance-configuration", "", "Model performance settings for the request.")
		bedrockAgentRuntime_invokeFlowCmd.Flags().Bool("no-enable-trace", false, "Specifies whether to return the trace for the flow or not.")
		bedrockAgentRuntime_invokeFlowCmd.MarkFlagRequired("flow-alias-identifier")
		bedrockAgentRuntime_invokeFlowCmd.MarkFlagRequired("flow-identifier")
		bedrockAgentRuntime_invokeFlowCmd.MarkFlagRequired("inputs")
		bedrockAgentRuntime_invokeFlowCmd.Flag("no-enable-trace").Hidden = true
	})
	bedrockAgentRuntimeCmd.AddCommand(bedrockAgentRuntime_invokeFlowCmd)
}
