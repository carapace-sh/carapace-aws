package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentcoreControl_updateMemoryCmd = &cobra.Command{
	Use:   "update-memory",
	Short: "Update an Amazon Bedrock AgentCore Memory resource memory.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentcoreControl_updateMemoryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgentcoreControl_updateMemoryCmd).Standalone()

		bedrockAgentcoreControl_updateMemoryCmd.Flags().String("client-token", "", "A client token is used for keeping track of idempotent requests.")
		bedrockAgentcoreControl_updateMemoryCmd.Flags().String("description", "", "The updated description of the AgentCore Memory resource.")
		bedrockAgentcoreControl_updateMemoryCmd.Flags().String("event-expiry-duration", "", "The number of days after which memory events will expire, between 7 and 365 days.")
		bedrockAgentcoreControl_updateMemoryCmd.Flags().String("memory-execution-role-arn", "", "The ARN of the IAM role that provides permissions for the AgentCore Memory resource.")
		bedrockAgentcoreControl_updateMemoryCmd.Flags().String("memory-id", "", "The unique identifier of the memory to update.")
		bedrockAgentcoreControl_updateMemoryCmd.Flags().String("memory-strategies", "", "The memory strategies to add, modify, or delete.")
		bedrockAgentcoreControl_updateMemoryCmd.MarkFlagRequired("memory-id")
	})
	bedrockAgentcoreControlCmd.AddCommand(bedrockAgentcoreControl_updateMemoryCmd)
}
