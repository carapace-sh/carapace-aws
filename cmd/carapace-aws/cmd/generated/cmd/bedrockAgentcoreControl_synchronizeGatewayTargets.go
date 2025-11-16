package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentcoreControl_synchronizeGatewayTargetsCmd = &cobra.Command{
	Use:   "synchronize-gateway-targets",
	Short: "The gateway targets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentcoreControl_synchronizeGatewayTargetsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgentcoreControl_synchronizeGatewayTargetsCmd).Standalone()

		bedrockAgentcoreControl_synchronizeGatewayTargetsCmd.Flags().String("gateway-identifier", "", "The gateway Identifier.")
		bedrockAgentcoreControl_synchronizeGatewayTargetsCmd.Flags().String("target-id-list", "", "The target ID list.")
		bedrockAgentcoreControl_synchronizeGatewayTargetsCmd.MarkFlagRequired("gateway-identifier")
		bedrockAgentcoreControl_synchronizeGatewayTargetsCmd.MarkFlagRequired("target-id-list")
	})
	bedrockAgentcoreControlCmd.AddCommand(bedrockAgentcoreControl_synchronizeGatewayTargetsCmd)
}
