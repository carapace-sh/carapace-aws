package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentcoreControl_createMemoryCmd = &cobra.Command{
	Use:   "create-memory",
	Short: "Creates a new Amazon Bedrock AgentCore Memory resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentcoreControl_createMemoryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgentcoreControl_createMemoryCmd).Standalone()

		bedrockAgentcoreControl_createMemoryCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier to ensure that the operation completes no more than one time.")
		bedrockAgentcoreControl_createMemoryCmd.Flags().String("description", "", "The description of the memory.")
		bedrockAgentcoreControl_createMemoryCmd.Flags().String("encryption-key-arn", "", "The Amazon Resource Name (ARN) of the KMS key used to encrypt the memory data.")
		bedrockAgentcoreControl_createMemoryCmd.Flags().String("event-expiry-duration", "", "The duration after which memory events expire.")
		bedrockAgentcoreControl_createMemoryCmd.Flags().String("memory-execution-role-arn", "", "The Amazon Resource Name (ARN) of the IAM role that provides permissions for the memory to access Amazon Web Services services.")
		bedrockAgentcoreControl_createMemoryCmd.Flags().String("memory-strategies", "", "The memory strategies to use for this memory.")
		bedrockAgentcoreControl_createMemoryCmd.Flags().String("name", "", "The name of the memory.")
		bedrockAgentcoreControl_createMemoryCmd.Flags().String("tags", "", "A map of tag keys and values to assign to an AgentCore Memory.")
		bedrockAgentcoreControl_createMemoryCmd.MarkFlagRequired("event-expiry-duration")
		bedrockAgentcoreControl_createMemoryCmd.MarkFlagRequired("name")
	})
	bedrockAgentcoreControlCmd.AddCommand(bedrockAgentcoreControl_createMemoryCmd)
}
