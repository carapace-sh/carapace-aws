package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentcoreControl_createGatewayCmd = &cobra.Command{
	Use:   "create-gateway",
	Short: "Creates a gateway for Amazon Bedrock Agent.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentcoreControl_createGatewayCmd).Standalone()

	bedrockAgentcoreControl_createGatewayCmd.Flags().String("authorizer-configuration", "", "The authorizer configuration for the gateway.")
	bedrockAgentcoreControl_createGatewayCmd.Flags().String("authorizer-type", "", "The type of authorizer to use for the gateway.")
	bedrockAgentcoreControl_createGatewayCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier to ensure that the API request completes no more than one time.")
	bedrockAgentcoreControl_createGatewayCmd.Flags().String("description", "", "The description of the gateway.")
	bedrockAgentcoreControl_createGatewayCmd.Flags().String("exception-level", "", "The level of detail in error messages returned when invoking the gateway.")
	bedrockAgentcoreControl_createGatewayCmd.Flags().String("kms-key-arn", "", "The Amazon Resource Name (ARN) of the KMS key used to encrypt data associated with the gateway.")
	bedrockAgentcoreControl_createGatewayCmd.Flags().String("name", "", "The name of the gateway.")
	bedrockAgentcoreControl_createGatewayCmd.Flags().String("protocol-configuration", "", "The configuration settings for the protocol specified in the `protocolType` parameter.")
	bedrockAgentcoreControl_createGatewayCmd.Flags().String("protocol-type", "", "The protocol type for the gateway.")
	bedrockAgentcoreControl_createGatewayCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of the IAM role that provides permissions for the gateway to access Amazon Web Services services.")
	bedrockAgentcoreControl_createGatewayCmd.Flags().String("tags", "", "A map of key-value pairs to associate with the gateway as metadata tags.")
	bedrockAgentcoreControl_createGatewayCmd.MarkFlagRequired("authorizer-type")
	bedrockAgentcoreControl_createGatewayCmd.MarkFlagRequired("name")
	bedrockAgentcoreControl_createGatewayCmd.MarkFlagRequired("protocol-type")
	bedrockAgentcoreControl_createGatewayCmd.MarkFlagRequired("role-arn")
	bedrockAgentcoreControlCmd.AddCommand(bedrockAgentcoreControl_createGatewayCmd)
}
