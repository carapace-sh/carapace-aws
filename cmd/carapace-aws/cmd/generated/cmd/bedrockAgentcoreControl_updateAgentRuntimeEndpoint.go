package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentcoreControl_updateAgentRuntimeEndpointCmd = &cobra.Command{
	Use:   "update-agent-runtime-endpoint",
	Short: "Updates an existing Amazon Bedrock AgentCore Runtime endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentcoreControl_updateAgentRuntimeEndpointCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgentcoreControl_updateAgentRuntimeEndpointCmd).Standalone()

		bedrockAgentcoreControl_updateAgentRuntimeEndpointCmd.Flags().String("agent-runtime-id", "", "The unique identifier of the AgentCore Runtime associated with the endpoint.")
		bedrockAgentcoreControl_updateAgentRuntimeEndpointCmd.Flags().String("agent-runtime-version", "", "The updated version of the AgentCore Runtime for the endpoint.")
		bedrockAgentcoreControl_updateAgentRuntimeEndpointCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier to ensure idempotency of the request.")
		bedrockAgentcoreControl_updateAgentRuntimeEndpointCmd.Flags().String("description", "", "The updated description of the AgentCore Runtime endpoint.")
		bedrockAgentcoreControl_updateAgentRuntimeEndpointCmd.Flags().String("endpoint-name", "", "The name of the AgentCore Runtime endpoint to update.")
		bedrockAgentcoreControl_updateAgentRuntimeEndpointCmd.MarkFlagRequired("agent-runtime-id")
		bedrockAgentcoreControl_updateAgentRuntimeEndpointCmd.MarkFlagRequired("endpoint-name")
	})
	bedrockAgentcoreControlCmd.AddCommand(bedrockAgentcoreControl_updateAgentRuntimeEndpointCmd)
}
