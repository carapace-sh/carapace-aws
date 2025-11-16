package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentcoreControl_getAgentRuntimeCmd = &cobra.Command{
	Use:   "get-agent-runtime",
	Short: "Gets an Amazon Bedrock AgentCore Runtime.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentcoreControl_getAgentRuntimeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgentcoreControl_getAgentRuntimeCmd).Standalone()

		bedrockAgentcoreControl_getAgentRuntimeCmd.Flags().String("agent-runtime-id", "", "The unique identifier of the AgentCore Runtime to retrieve.")
		bedrockAgentcoreControl_getAgentRuntimeCmd.Flags().String("agent-runtime-version", "", "The version of the AgentCore Runtime to retrieve.")
		bedrockAgentcoreControl_getAgentRuntimeCmd.MarkFlagRequired("agent-runtime-id")
	})
	bedrockAgentcoreControlCmd.AddCommand(bedrockAgentcoreControl_getAgentRuntimeCmd)
}
