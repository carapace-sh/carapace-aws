package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentcoreControl_createAgentRuntimeEndpointCmd = &cobra.Command{
	Use:   "create-agent-runtime-endpoint",
	Short: "Creates an AgentCore Runtime endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentcoreControl_createAgentRuntimeEndpointCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgentcoreControl_createAgentRuntimeEndpointCmd).Standalone()

		bedrockAgentcoreControl_createAgentRuntimeEndpointCmd.Flags().String("agent-runtime-id", "", "The unique identifier of the AgentCore Runtime to create an endpoint for.")
		bedrockAgentcoreControl_createAgentRuntimeEndpointCmd.Flags().String("agent-runtime-version", "", "The version of the AgentCore Runtime to use for the endpoint.")
		bedrockAgentcoreControl_createAgentRuntimeEndpointCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier to ensure idempotency of the request.")
		bedrockAgentcoreControl_createAgentRuntimeEndpointCmd.Flags().String("description", "", "The description of the AgentCore Runtime endpoint.")
		bedrockAgentcoreControl_createAgentRuntimeEndpointCmd.Flags().String("name", "", "The name of the AgentCore Runtime endpoint.")
		bedrockAgentcoreControl_createAgentRuntimeEndpointCmd.Flags().String("tags", "", "A map of tag keys and values to assign to the agent runtime endpoint.")
		bedrockAgentcoreControl_createAgentRuntimeEndpointCmd.MarkFlagRequired("agent-runtime-id")
		bedrockAgentcoreControl_createAgentRuntimeEndpointCmd.MarkFlagRequired("name")
	})
	bedrockAgentcoreControlCmd.AddCommand(bedrockAgentcoreControl_createAgentRuntimeEndpointCmd)
}
