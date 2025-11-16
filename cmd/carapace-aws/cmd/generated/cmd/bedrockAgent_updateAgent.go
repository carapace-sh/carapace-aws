package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgent_updateAgentCmd = &cobra.Command{
	Use:   "update-agent",
	Short: "Updates the configuration of an agent.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgent_updateAgentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgent_updateAgentCmd).Standalone()

		bedrockAgent_updateAgentCmd.Flags().String("agent-collaboration", "", "The agent's collaboration role.")
		bedrockAgent_updateAgentCmd.Flags().String("agent-id", "", "The unique identifier of the agent.")
		bedrockAgent_updateAgentCmd.Flags().String("agent-name", "", "Specifies a new name for the agent.")
		bedrockAgent_updateAgentCmd.Flags().String("agent-resource-role-arn", "", "The Amazon Resource Name (ARN) of the IAM role with permissions to invoke API operations on the agent.")
		bedrockAgent_updateAgentCmd.Flags().String("custom-orchestration", "", "Contains details of the custom orchestration configured for the agent.")
		bedrockAgent_updateAgentCmd.Flags().String("customer-encryption-key-arn", "", "The Amazon Resource Name (ARN) of the KMS key with which to encrypt the agent.")
		bedrockAgent_updateAgentCmd.Flags().String("description", "", "Specifies a new description of the agent.")
		bedrockAgent_updateAgentCmd.Flags().String("foundation-model", "", "The identifier for the model that you want to be used for orchestration by the agent you create.")
		bedrockAgent_updateAgentCmd.Flags().String("guardrail-configuration", "", "The unique Guardrail configuration assigned to the agent when it is updated.")
		bedrockAgent_updateAgentCmd.Flags().String("idle-session-ttlin-seconds", "", "The number of seconds for which Amazon Bedrock keeps information about a user's conversation with the agent.")
		bedrockAgent_updateAgentCmd.Flags().String("instruction", "", "Specifies new instructions that tell the agent what it should do and how it should interact with users.")
		bedrockAgent_updateAgentCmd.Flags().String("memory-configuration", "", "Specifies the new memory configuration for the agent.")
		bedrockAgent_updateAgentCmd.Flags().String("orchestration-type", "", "Specifies the type of orchestration strategy for the agent.")
		bedrockAgent_updateAgentCmd.Flags().String("prompt-override-configuration", "", "Contains configurations to override prompts in different parts of an agent sequence.")
		bedrockAgent_updateAgentCmd.MarkFlagRequired("agent-id")
		bedrockAgent_updateAgentCmd.MarkFlagRequired("agent-name")
		bedrockAgent_updateAgentCmd.MarkFlagRequired("agent-resource-role-arn")
		bedrockAgent_updateAgentCmd.MarkFlagRequired("foundation-model")
	})
	bedrockAgentCmd.AddCommand(bedrockAgent_updateAgentCmd)
}
