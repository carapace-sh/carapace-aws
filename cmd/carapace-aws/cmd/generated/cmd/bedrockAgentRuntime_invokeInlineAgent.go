package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentRuntime_invokeInlineAgentCmd = &cobra.Command{
	Use:   "invoke-inline-agent",
	Short: "Invokes an inline Amazon Bedrock agent using the configurations you provide with the request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentRuntime_invokeInlineAgentCmd).Standalone()

	bedrockAgentRuntime_invokeInlineAgentCmd.Flags().String("action-groups", "", "A list of action groups with each action group defining the action the inline agent needs to carry out.")
	bedrockAgentRuntime_invokeInlineAgentCmd.Flags().String("agent-collaboration", "", "Defines how the inline collaborator agent handles information across multiple collaborator agents to coordinate a final response.")
	bedrockAgentRuntime_invokeInlineAgentCmd.Flags().String("agent-name", "", "The name for the agent.")
	bedrockAgentRuntime_invokeInlineAgentCmd.Flags().String("bedrock-model-configurations", "", "Model settings for the request.")
	bedrockAgentRuntime_invokeInlineAgentCmd.Flags().String("collaborator-configurations", "", "Settings for an inline agent collaborator called with [InvokeInlineAgent](https://docs.aws.amazon.com/bedrock/latest/APIReference/API_agent-runtime_InvokeInlineAgent.html).")
	bedrockAgentRuntime_invokeInlineAgentCmd.Flags().String("collaborators", "", "List of collaborator inline agents.")
	bedrockAgentRuntime_invokeInlineAgentCmd.Flags().String("custom-orchestration", "", "Contains details of the custom orchestration configured for the agent.")
	bedrockAgentRuntime_invokeInlineAgentCmd.Flags().String("customer-encryption-key-arn", "", "The Amazon Resource Name (ARN) of the Amazon Web Services KMS key to use to encrypt your inline agent.")
	bedrockAgentRuntime_invokeInlineAgentCmd.Flags().Bool("enable-trace", false, "Specifies whether to turn on the trace or not to track the agent's reasoning process.")
	bedrockAgentRuntime_invokeInlineAgentCmd.Flags().Bool("end-session", false, "Specifies whether to end the session with the inline agent or not.")
	bedrockAgentRuntime_invokeInlineAgentCmd.Flags().String("foundation-model", "", "The [model identifier (ID)](https://docs.aws.amazon.com/bedrock/latest/userguide/model-ids.html#model-ids-arns) of the model to use for orchestration by the inline agent.")
	bedrockAgentRuntime_invokeInlineAgentCmd.Flags().String("guardrail-configuration", "", "The [guardrails](https://docs.aws.amazon.com/bedrock/latest/userguide/guardrails.html) to assign to the inline agent.")
	bedrockAgentRuntime_invokeInlineAgentCmd.Flags().String("idle-session-ttlin-seconds", "", "The number of seconds for which the inline agent should maintain session information.")
	bedrockAgentRuntime_invokeInlineAgentCmd.Flags().String("inline-session-state", "", "Parameters that specify the various attributes of a sessions.")
	bedrockAgentRuntime_invokeInlineAgentCmd.Flags().String("input-text", "", "The prompt text to send to the agent.")
	bedrockAgentRuntime_invokeInlineAgentCmd.Flags().String("instruction", "", "The instructions that tell the inline agent what it should do and how it should interact with users.")
	bedrockAgentRuntime_invokeInlineAgentCmd.Flags().String("knowledge-bases", "", "Contains information of the knowledge bases to associate with.")
	bedrockAgentRuntime_invokeInlineAgentCmd.Flags().Bool("no-enable-trace", false, "Specifies whether to turn on the trace or not to track the agent's reasoning process.")
	bedrockAgentRuntime_invokeInlineAgentCmd.Flags().Bool("no-end-session", false, "Specifies whether to end the session with the inline agent or not.")
	bedrockAgentRuntime_invokeInlineAgentCmd.Flags().String("orchestration-type", "", "Specifies the type of orchestration strategy for the agent.")
	bedrockAgentRuntime_invokeInlineAgentCmd.Flags().String("prompt-creation-configurations", "", "Specifies parameters that control how the service populates the agent prompt for an `InvokeInlineAgent` request.")
	bedrockAgentRuntime_invokeInlineAgentCmd.Flags().String("prompt-override-configuration", "", "Configurations for advanced prompts used to override the default prompts to enhance the accuracy of the inline agent.")
	bedrockAgentRuntime_invokeInlineAgentCmd.Flags().String("session-id", "", "The unique identifier of the session.")
	bedrockAgentRuntime_invokeInlineAgentCmd.Flags().String("streaming-configurations", "", "Specifies the configurations for streaming.")
	bedrockAgentRuntime_invokeInlineAgentCmd.MarkFlagRequired("foundation-model")
	bedrockAgentRuntime_invokeInlineAgentCmd.MarkFlagRequired("instruction")
	bedrockAgentRuntime_invokeInlineAgentCmd.Flag("no-enable-trace").Hidden = true
	bedrockAgentRuntime_invokeInlineAgentCmd.Flag("no-end-session").Hidden = true
	bedrockAgentRuntime_invokeInlineAgentCmd.MarkFlagRequired("session-id")
	bedrockAgentRuntimeCmd.AddCommand(bedrockAgentRuntime_invokeInlineAgentCmd)
}
