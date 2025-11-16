package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentRuntime_invokeAgentCmd = &cobra.Command{
	Use:   "invoke-agent",
	Short: "Sends a prompt for the agent to process and respond to.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentRuntime_invokeAgentCmd).Standalone()

	bedrockAgentRuntime_invokeAgentCmd.Flags().String("agent-alias-id", "", "The alias of the agent to use.")
	bedrockAgentRuntime_invokeAgentCmd.Flags().String("agent-id", "", "The unique identifier of the agent to use.")
	bedrockAgentRuntime_invokeAgentCmd.Flags().String("bedrock-model-configurations", "", "Model performance settings for the request.")
	bedrockAgentRuntime_invokeAgentCmd.Flags().Bool("enable-trace", false, "Specifies whether to turn on the trace or not to track the agent's reasoning process.")
	bedrockAgentRuntime_invokeAgentCmd.Flags().Bool("end-session", false, "Specifies whether to end the session with the agent or not.")
	bedrockAgentRuntime_invokeAgentCmd.Flags().String("input-text", "", "The prompt text to send the agent.")
	bedrockAgentRuntime_invokeAgentCmd.Flags().String("memory-id", "", "The unique identifier of the agent memory.")
	bedrockAgentRuntime_invokeAgentCmd.Flags().Bool("no-enable-trace", false, "Specifies whether to turn on the trace or not to track the agent's reasoning process.")
	bedrockAgentRuntime_invokeAgentCmd.Flags().Bool("no-end-session", false, "Specifies whether to end the session with the agent or not.")
	bedrockAgentRuntime_invokeAgentCmd.Flags().String("prompt-creation-configurations", "", "Specifies parameters that control how the service populates the agent prompt for an `InvokeAgent` request.")
	bedrockAgentRuntime_invokeAgentCmd.Flags().String("session-id", "", "The unique identifier of the session.")
	bedrockAgentRuntime_invokeAgentCmd.Flags().String("session-state", "", "Contains parameters that specify various attributes of the session.")
	bedrockAgentRuntime_invokeAgentCmd.Flags().String("source-arn", "", "The ARN of the resource making the request.")
	bedrockAgentRuntime_invokeAgentCmd.Flags().String("streaming-configurations", "", "Specifies the configurations for streaming.")
	bedrockAgentRuntime_invokeAgentCmd.MarkFlagRequired("agent-alias-id")
	bedrockAgentRuntime_invokeAgentCmd.MarkFlagRequired("agent-id")
	bedrockAgentRuntime_invokeAgentCmd.Flag("no-enable-trace").Hidden = true
	bedrockAgentRuntime_invokeAgentCmd.Flag("no-end-session").Hidden = true
	bedrockAgentRuntime_invokeAgentCmd.MarkFlagRequired("session-id")
	bedrockAgentRuntimeCmd.AddCommand(bedrockAgentRuntime_invokeAgentCmd)
}
