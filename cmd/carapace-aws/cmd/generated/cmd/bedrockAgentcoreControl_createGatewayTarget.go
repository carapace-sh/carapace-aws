package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentcoreControl_createGatewayTargetCmd = &cobra.Command{
	Use:   "create-gateway-target",
	Short: "Creates a target for a gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentcoreControl_createGatewayTargetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgentcoreControl_createGatewayTargetCmd).Standalone()

		bedrockAgentcoreControl_createGatewayTargetCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier to ensure that the API request completes no more than one time.")
		bedrockAgentcoreControl_createGatewayTargetCmd.Flags().String("credential-provider-configurations", "", "The credential provider configurations for the target.")
		bedrockAgentcoreControl_createGatewayTargetCmd.Flags().String("description", "", "The description of the gateway target.")
		bedrockAgentcoreControl_createGatewayTargetCmd.Flags().String("gateway-identifier", "", "The identifier of the gateway to create a target for.")
		bedrockAgentcoreControl_createGatewayTargetCmd.Flags().String("name", "", "The name of the gateway target.")
		bedrockAgentcoreControl_createGatewayTargetCmd.Flags().String("target-configuration", "", "The configuration settings for the target, including endpoint information and schema definitions.")
		bedrockAgentcoreControl_createGatewayTargetCmd.MarkFlagRequired("gateway-identifier")
		bedrockAgentcoreControl_createGatewayTargetCmd.MarkFlagRequired("name")
		bedrockAgentcoreControl_createGatewayTargetCmd.MarkFlagRequired("target-configuration")
	})
	bedrockAgentcoreControlCmd.AddCommand(bedrockAgentcoreControl_createGatewayTargetCmd)
}
