package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgent_createAgentActionGroupCmd = &cobra.Command{
	Use:   "create-agent-action-group",
	Short: "Creates an action group for an agent.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgent_createAgentActionGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgent_createAgentActionGroupCmd).Standalone()

		bedrockAgent_createAgentActionGroupCmd.Flags().String("action-group-executor", "", "The Amazon Resource Name (ARN) of the Lambda function containing the business logic that is carried out upon invoking the action or the custom control method for handling the information elicited from the user.")
		bedrockAgent_createAgentActionGroupCmd.Flags().String("action-group-name", "", "The name to give the action group.")
		bedrockAgent_createAgentActionGroupCmd.Flags().String("action-group-state", "", "Specifies whether the action group is available for the agent to invoke or not when sending an [InvokeAgent](https://docs.aws.amazon.com/bedrock/latest/APIReference/API_agent-runtime_InvokeAgent.html) request.")
		bedrockAgent_createAgentActionGroupCmd.Flags().String("agent-id", "", "The unique identifier of the agent for which to create the action group.")
		bedrockAgent_createAgentActionGroupCmd.Flags().String("agent-version", "", "The version of the agent for which to create the action group.")
		bedrockAgent_createAgentActionGroupCmd.Flags().String("api-schema", "", "Contains either details about the S3 object containing the OpenAPI schema for the action group or the JSON or YAML-formatted payload defining the schema.")
		bedrockAgent_createAgentActionGroupCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier to ensure that the API request completes no more than one time.")
		bedrockAgent_createAgentActionGroupCmd.Flags().String("description", "", "A description of the action group.")
		bedrockAgent_createAgentActionGroupCmd.Flags().String("function-schema", "", "Contains details about the function schema for the action group or the JSON or YAML-formatted payload defining the schema.")
		bedrockAgent_createAgentActionGroupCmd.Flags().String("parent-action-group-signature", "", "Specify a built-in or computer use action for this action group.")
		bedrockAgent_createAgentActionGroupCmd.Flags().String("parent-action-group-signature-params", "", "The configuration settings for a computer use action.")
		bedrockAgent_createAgentActionGroupCmd.MarkFlagRequired("action-group-name")
		bedrockAgent_createAgentActionGroupCmd.MarkFlagRequired("agent-id")
		bedrockAgent_createAgentActionGroupCmd.MarkFlagRequired("agent-version")
	})
	bedrockAgentCmd.AddCommand(bedrockAgent_createAgentActionGroupCmd)
}
