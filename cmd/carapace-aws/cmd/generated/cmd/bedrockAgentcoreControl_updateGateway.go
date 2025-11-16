package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentcoreControl_updateGatewayCmd = &cobra.Command{
	Use:   "update-gateway",
	Short: "Updates an existing gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentcoreControl_updateGatewayCmd).Standalone()

	bedrockAgentcoreControl_updateGatewayCmd.Flags().String("authorizer-configuration", "", "The updated authorizer configuration for the gateway.")
	bedrockAgentcoreControl_updateGatewayCmd.Flags().String("authorizer-type", "", "The updated authorizer type for the gateway.")
	bedrockAgentcoreControl_updateGatewayCmd.Flags().String("description", "", "The updated description for the gateway.")
	bedrockAgentcoreControl_updateGatewayCmd.Flags().String("exception-level", "", "The level of detail in error messages returned when invoking the gateway.")
	bedrockAgentcoreControl_updateGatewayCmd.Flags().String("gateway-identifier", "", "The identifier of the gateway to update.")
	bedrockAgentcoreControl_updateGatewayCmd.Flags().String("kms-key-arn", "", "The updated ARN of the KMS key used to encrypt the gateway.")
	bedrockAgentcoreControl_updateGatewayCmd.Flags().String("name", "", "The name of the gateway.")
	bedrockAgentcoreControl_updateGatewayCmd.Flags().String("protocol-configuration", "", "")
	bedrockAgentcoreControl_updateGatewayCmd.Flags().String("protocol-type", "", "The updated protocol type for the gateway.")
	bedrockAgentcoreControl_updateGatewayCmd.Flags().String("role-arn", "", "The updated IAM role ARN that provides permissions for the gateway.")
	bedrockAgentcoreControl_updateGatewayCmd.MarkFlagRequired("authorizer-type")
	bedrockAgentcoreControl_updateGatewayCmd.MarkFlagRequired("gateway-identifier")
	bedrockAgentcoreControl_updateGatewayCmd.MarkFlagRequired("name")
	bedrockAgentcoreControl_updateGatewayCmd.MarkFlagRequired("protocol-type")
	bedrockAgentcoreControl_updateGatewayCmd.MarkFlagRequired("role-arn")
	bedrockAgentcoreControlCmd.AddCommand(bedrockAgentcoreControl_updateGatewayCmd)
}
