package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentcoreControl_updateGatewayTargetCmd = &cobra.Command{
	Use:   "update-gateway-target",
	Short: "Updates an existing gateway target.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentcoreControl_updateGatewayTargetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgentcoreControl_updateGatewayTargetCmd).Standalone()

		bedrockAgentcoreControl_updateGatewayTargetCmd.Flags().String("credential-provider-configurations", "", "The updated credential provider configurations for the gateway target.")
		bedrockAgentcoreControl_updateGatewayTargetCmd.Flags().String("description", "", "The updated description for the gateway target.")
		bedrockAgentcoreControl_updateGatewayTargetCmd.Flags().String("gateway-identifier", "", "The unique identifier of the gateway associated with the target.")
		bedrockAgentcoreControl_updateGatewayTargetCmd.Flags().String("name", "", "The updated name for the gateway target.")
		bedrockAgentcoreControl_updateGatewayTargetCmd.Flags().String("target-configuration", "", "")
		bedrockAgentcoreControl_updateGatewayTargetCmd.Flags().String("target-id", "", "The unique identifier of the gateway target to update.")
		bedrockAgentcoreControl_updateGatewayTargetCmd.MarkFlagRequired("gateway-identifier")
		bedrockAgentcoreControl_updateGatewayTargetCmd.MarkFlagRequired("name")
		bedrockAgentcoreControl_updateGatewayTargetCmd.MarkFlagRequired("target-configuration")
		bedrockAgentcoreControl_updateGatewayTargetCmd.MarkFlagRequired("target-id")
	})
	bedrockAgentcoreControlCmd.AddCommand(bedrockAgentcoreControl_updateGatewayTargetCmd)
}
