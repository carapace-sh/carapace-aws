package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentcoreControl_updateAgentRuntimeCmd = &cobra.Command{
	Use:   "update-agent-runtime",
	Short: "Updates an existing Amazon Secure Agent.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentcoreControl_updateAgentRuntimeCmd).Standalone()

	bedrockAgentcoreControl_updateAgentRuntimeCmd.Flags().String("agent-runtime-artifact", "", "The updated artifact of the AgentCore Runtime.")
	bedrockAgentcoreControl_updateAgentRuntimeCmd.Flags().String("agent-runtime-id", "", "The unique identifier of the AgentCore Runtime to update.")
	bedrockAgentcoreControl_updateAgentRuntimeCmd.Flags().String("authorizer-configuration", "", "The updated authorizer configuration for the AgentCore Runtime.")
	bedrockAgentcoreControl_updateAgentRuntimeCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier to ensure idempotency of the request.")
	bedrockAgentcoreControl_updateAgentRuntimeCmd.Flags().String("description", "", "The updated description of the AgentCore Runtime.")
	bedrockAgentcoreControl_updateAgentRuntimeCmd.Flags().String("environment-variables", "", "Updated environment variables to set in the AgentCore Runtime environment.")
	bedrockAgentcoreControl_updateAgentRuntimeCmd.Flags().String("lifecycle-configuration", "", "The updated life cycle configuration for the AgentCore Runtime.")
	bedrockAgentcoreControl_updateAgentRuntimeCmd.Flags().String("network-configuration", "", "The updated network configuration for the AgentCore Runtime.")
	bedrockAgentcoreControl_updateAgentRuntimeCmd.Flags().String("protocol-configuration", "", "")
	bedrockAgentcoreControl_updateAgentRuntimeCmd.Flags().String("request-header-configuration", "", "The updated configuration for HTTP request headers that will be passed through to the runtime.")
	bedrockAgentcoreControl_updateAgentRuntimeCmd.Flags().String("role-arn", "", "The updated IAM role ARN that provides permissions for the AgentCore Runtime.")
	bedrockAgentcoreControl_updateAgentRuntimeCmd.MarkFlagRequired("agent-runtime-artifact")
	bedrockAgentcoreControl_updateAgentRuntimeCmd.MarkFlagRequired("agent-runtime-id")
	bedrockAgentcoreControl_updateAgentRuntimeCmd.MarkFlagRequired("network-configuration")
	bedrockAgentcoreControl_updateAgentRuntimeCmd.MarkFlagRequired("role-arn")
	bedrockAgentcoreControlCmd.AddCommand(bedrockAgentcoreControl_updateAgentRuntimeCmd)
}
