package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgent_createAgentCmd = &cobra.Command{
	Use:   "create-agent",
	Short: "Creates an agent that orchestrates interactions between foundation models, data sources, software applications, user conversations, and APIs to carry out tasks to help customers.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgent_createAgentCmd).Standalone()

	bedrockAgent_createAgentCmd.Flags().String("agent-collaboration", "", "The agent's collaboration role.")
	bedrockAgent_createAgentCmd.Flags().String("agent-name", "", "A name for the agent that you create.")
	bedrockAgent_createAgentCmd.Flags().String("agent-resource-role-arn", "", "The Amazon Resource Name (ARN) of the IAM role with permissions to invoke API operations on the agent.")
	bedrockAgent_createAgentCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier to ensure that the API request completes no more than one time.")
	bedrockAgent_createAgentCmd.Flags().String("custom-orchestration", "", "Contains details of the custom orchestration configured for the agent.")
	bedrockAgent_createAgentCmd.Flags().String("customer-encryption-key-arn", "", "The Amazon Resource Name (ARN) of the KMS key with which to encrypt the agent.")
	bedrockAgent_createAgentCmd.Flags().String("description", "", "A description of the agent.")
	bedrockAgent_createAgentCmd.Flags().String("foundation-model", "", "The identifier for the model that you want to be used for orchestration by the agent you create.")
	bedrockAgent_createAgentCmd.Flags().String("guardrail-configuration", "", "The unique Guardrail configuration assigned to the agent when it is created.")
	bedrockAgent_createAgentCmd.Flags().String("idle-session-ttlin-seconds", "", "The number of seconds for which Amazon Bedrock keeps information about a user's conversation with the agent.")
	bedrockAgent_createAgentCmd.Flags().String("instruction", "", "Instructions that tell the agent what it should do and how it should interact with users.")
	bedrockAgent_createAgentCmd.Flags().String("memory-configuration", "", "Contains the details of the memory configured for the agent.")
	bedrockAgent_createAgentCmd.Flags().String("orchestration-type", "", "Specifies the type of orchestration strategy for the agent.")
	bedrockAgent_createAgentCmd.Flags().String("prompt-override-configuration", "", "Contains configurations to override prompts in different parts of an agent sequence.")
	bedrockAgent_createAgentCmd.Flags().String("tags", "", "Any tags that you want to attach to the agent.")
	bedrockAgent_createAgentCmd.MarkFlagRequired("agent-name")
	bedrockAgentCmd.AddCommand(bedrockAgent_createAgentCmd)
}
