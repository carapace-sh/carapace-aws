package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentcoreControl_getGatewayTargetCmd = &cobra.Command{
	Use:   "get-gateway-target",
	Short: "Retrieves information about a specific gateway target.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentcoreControl_getGatewayTargetCmd).Standalone()

	bedrockAgentcoreControl_getGatewayTargetCmd.Flags().String("gateway-identifier", "", "The identifier of the gateway that contains the target.")
	bedrockAgentcoreControl_getGatewayTargetCmd.Flags().String("target-id", "", "The unique identifier of the target to retrieve.")
	bedrockAgentcoreControl_getGatewayTargetCmd.MarkFlagRequired("gateway-identifier")
	bedrockAgentcoreControl_getGatewayTargetCmd.MarkFlagRequired("target-id")
	bedrockAgentcoreControlCmd.AddCommand(bedrockAgentcoreControl_getGatewayTargetCmd)
}
