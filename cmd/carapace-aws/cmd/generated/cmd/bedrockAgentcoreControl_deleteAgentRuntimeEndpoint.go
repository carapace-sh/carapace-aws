package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentcoreControl_deleteAgentRuntimeEndpointCmd = &cobra.Command{
	Use:   "delete-agent-runtime-endpoint",
	Short: "Deletes an AAgentCore Runtime endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentcoreControl_deleteAgentRuntimeEndpointCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgentcoreControl_deleteAgentRuntimeEndpointCmd).Standalone()

		bedrockAgentcoreControl_deleteAgentRuntimeEndpointCmd.Flags().String("agent-runtime-id", "", "The unique identifier of the AgentCore Runtime associated with the endpoint.")
		bedrockAgentcoreControl_deleteAgentRuntimeEndpointCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier to ensure idempotency of the request.")
		bedrockAgentcoreControl_deleteAgentRuntimeEndpointCmd.Flags().String("endpoint-name", "", "The name of the AgentCore Runtime endpoint to delete.")
		bedrockAgentcoreControl_deleteAgentRuntimeEndpointCmd.MarkFlagRequired("agent-runtime-id")
		bedrockAgentcoreControl_deleteAgentRuntimeEndpointCmd.MarkFlagRequired("endpoint-name")
	})
	bedrockAgentcoreControlCmd.AddCommand(bedrockAgentcoreControl_deleteAgentRuntimeEndpointCmd)
}
