package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgent_updateAgentActionGroupCmd = &cobra.Command{
	Use:   "update-agent-action-group",
	Short: "Updates the configuration for an action group for an agent.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgent_updateAgentActionGroupCmd).Standalone()

	bedrockAgent_updateAgentActionGroupCmd.Flags().String("action-group-executor", "", "The Amazon Resource Name (ARN) of the Lambda function containing the business logic that is carried out upon invoking the action.")
	bedrockAgent_updateAgentActionGroupCmd.Flags().String("action-group-id", "", "The unique identifier of the action group.")
	bedrockAgent_updateAgentActionGroupCmd.Flags().String("action-group-name", "", "Specifies a new name for the action group.")
	bedrockAgent_updateAgentActionGroupCmd.Flags().String("action-group-state", "", "Specifies whether the action group is available for the agent to invoke or not when sending an [InvokeAgent](https://docs.aws.amazon.com/bedrock/latest/APIReference/API_agent-runtime_InvokeAgent.html) request.")
	bedrockAgent_updateAgentActionGroupCmd.Flags().String("agent-id", "", "The unique identifier of the agent for which to update the action group.")
	bedrockAgent_updateAgentActionGroupCmd.Flags().String("agent-version", "", "The unique identifier of the agent version for which to update the action group.")
	bedrockAgent_updateAgentActionGroupCmd.Flags().String("api-schema", "", "Contains either details about the S3 object containing the OpenAPI schema for the action group or the JSON or YAML-formatted payload defining the schema.")
	bedrockAgent_updateAgentActionGroupCmd.Flags().String("description", "", "Specifies a new name for the action group.")
	bedrockAgent_updateAgentActionGroupCmd.Flags().String("function-schema", "", "Contains details about the function schema for the action group or the JSON or YAML-formatted payload defining the schema.")
	bedrockAgent_updateAgentActionGroupCmd.Flags().String("parent-action-group-signature", "", "Update the built-in or computer use action for this action group.")
	bedrockAgent_updateAgentActionGroupCmd.Flags().String("parent-action-group-signature-params", "", "The configuration settings for a computer use action.")
	bedrockAgent_updateAgentActionGroupCmd.MarkFlagRequired("action-group-id")
	bedrockAgent_updateAgentActionGroupCmd.MarkFlagRequired("action-group-name")
	bedrockAgent_updateAgentActionGroupCmd.MarkFlagRequired("agent-id")
	bedrockAgent_updateAgentActionGroupCmd.MarkFlagRequired("agent-version")
	bedrockAgentCmd.AddCommand(bedrockAgent_updateAgentActionGroupCmd)
}
