package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentcoreControl_createAgentRuntimeCmd = &cobra.Command{
	Use:   "create-agent-runtime",
	Short: "Creates an Amazon Bedrock AgentCore Runtime.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentcoreControl_createAgentRuntimeCmd).Standalone()

	bedrockAgentcoreControl_createAgentRuntimeCmd.Flags().String("agent-runtime-artifact", "", "The artifact of the AgentCore Runtime.")
	bedrockAgentcoreControl_createAgentRuntimeCmd.Flags().String("agent-runtime-name", "", "The name of the AgentCore Runtime.")
	bedrockAgentcoreControl_createAgentRuntimeCmd.Flags().String("authorizer-configuration", "", "The authorizer configuration for the AgentCore Runtime.")
	bedrockAgentcoreControl_createAgentRuntimeCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier to ensure idempotency of the request.")
	bedrockAgentcoreControl_createAgentRuntimeCmd.Flags().String("description", "", "The description of the AgentCore Runtime.")
	bedrockAgentcoreControl_createAgentRuntimeCmd.Flags().String("environment-variables", "", "Environment variables to set in the AgentCore Runtime environment.")
	bedrockAgentcoreControl_createAgentRuntimeCmd.Flags().String("lifecycle-configuration", "", "The life cycle configuration for the AgentCore Runtime.")
	bedrockAgentcoreControl_createAgentRuntimeCmd.Flags().String("network-configuration", "", "The network configuration for the AgentCore Runtime.")
	bedrockAgentcoreControl_createAgentRuntimeCmd.Flags().String("protocol-configuration", "", "")
	bedrockAgentcoreControl_createAgentRuntimeCmd.Flags().String("request-header-configuration", "", "Configuration for HTTP request headers that will be passed through to the runtime.")
	bedrockAgentcoreControl_createAgentRuntimeCmd.Flags().String("role-arn", "", "The IAM role ARN that provides permissions for the AgentCore Runtime.")
	bedrockAgentcoreControl_createAgentRuntimeCmd.Flags().String("tags", "", "A map of tag keys and values to assign to the agent runtime.")
	bedrockAgentcoreControl_createAgentRuntimeCmd.MarkFlagRequired("agent-runtime-artifact")
	bedrockAgentcoreControl_createAgentRuntimeCmd.MarkFlagRequired("agent-runtime-name")
	bedrockAgentcoreControl_createAgentRuntimeCmd.MarkFlagRequired("network-configuration")
	bedrockAgentcoreControl_createAgentRuntimeCmd.MarkFlagRequired("role-arn")
	bedrockAgentcoreControlCmd.AddCommand(bedrockAgentcoreControl_createAgentRuntimeCmd)
}
