package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentcore_invokeAgentRuntimeCmd = &cobra.Command{
	Use:   "invoke-agent-runtime",
	Short: "Sends a request to an agent or tool hosted in an Amazon Bedrock AgentCore Runtime and receives responses in real-time.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentcore_invokeAgentRuntimeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgentcore_invokeAgentRuntimeCmd).Standalone()

		bedrockAgentcore_invokeAgentRuntimeCmd.Flags().String("accept", "", "The desired MIME type for the response from the agent runtime.")
		bedrockAgentcore_invokeAgentRuntimeCmd.Flags().String("account-id", "", "The identifier of the Amazon Web Services account for the agent runtime resource.")
		bedrockAgentcore_invokeAgentRuntimeCmd.Flags().String("agent-runtime-arn", "", "The Amazon Web Services Resource Name (ARN) of the agent runtime to invoke.")
		bedrockAgentcore_invokeAgentRuntimeCmd.Flags().String("baggage", "", "Additional context information for distributed tracing.")
		bedrockAgentcore_invokeAgentRuntimeCmd.Flags().String("content-type", "", "The MIME type of the input data in the payload.")
		bedrockAgentcore_invokeAgentRuntimeCmd.Flags().String("mcp-protocol-version", "", "The version of the MCP protocol being used.")
		bedrockAgentcore_invokeAgentRuntimeCmd.Flags().String("mcp-session-id", "", "The identifier of the MCP session.")
		bedrockAgentcore_invokeAgentRuntimeCmd.Flags().String("payload", "", "The input data to send to the agent runtime.")
		bedrockAgentcore_invokeAgentRuntimeCmd.Flags().String("qualifier", "", "The qualifier to use for the agent runtime.")
		bedrockAgentcore_invokeAgentRuntimeCmd.Flags().String("runtime-session-id", "", "The identifier of the runtime session.")
		bedrockAgentcore_invokeAgentRuntimeCmd.Flags().String("runtime-user-id", "", "The identifier of the runtime user.")
		bedrockAgentcore_invokeAgentRuntimeCmd.Flags().String("trace-id", "", "The trace identifier for request tracking.")
		bedrockAgentcore_invokeAgentRuntimeCmd.Flags().String("trace-parent", "", "The parent trace information for distributed tracing.")
		bedrockAgentcore_invokeAgentRuntimeCmd.Flags().String("trace-state", "", "The trace state information for distributed tracing.")
		bedrockAgentcore_invokeAgentRuntimeCmd.MarkFlagRequired("agent-runtime-arn")
		bedrockAgentcore_invokeAgentRuntimeCmd.MarkFlagRequired("payload")
	})
	bedrockAgentcoreCmd.AddCommand(bedrockAgentcore_invokeAgentRuntimeCmd)
}
