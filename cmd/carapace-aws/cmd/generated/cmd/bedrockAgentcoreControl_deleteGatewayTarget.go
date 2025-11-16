package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentcoreControl_deleteGatewayTargetCmd = &cobra.Command{
	Use:   "delete-gateway-target",
	Short: "Deletes a gateway target.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentcoreControl_deleteGatewayTargetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgentcoreControl_deleteGatewayTargetCmd).Standalone()

		bedrockAgentcoreControl_deleteGatewayTargetCmd.Flags().String("gateway-identifier", "", "The unique identifier of the gateway associated with the target.")
		bedrockAgentcoreControl_deleteGatewayTargetCmd.Flags().String("target-id", "", "The unique identifier of the gateway target to delete.")
		bedrockAgentcoreControl_deleteGatewayTargetCmd.MarkFlagRequired("gateway-identifier")
		bedrockAgentcoreControl_deleteGatewayTargetCmd.MarkFlagRequired("target-id")
	})
	bedrockAgentcoreControlCmd.AddCommand(bedrockAgentcoreControl_deleteGatewayTargetCmd)
}
