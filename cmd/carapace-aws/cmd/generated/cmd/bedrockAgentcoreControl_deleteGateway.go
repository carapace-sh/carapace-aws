package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentcoreControl_deleteGatewayCmd = &cobra.Command{
	Use:   "delete-gateway",
	Short: "Deletes a gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentcoreControl_deleteGatewayCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgentcoreControl_deleteGatewayCmd).Standalone()

		bedrockAgentcoreControl_deleteGatewayCmd.Flags().String("gateway-identifier", "", "The identifier of the gateway to delete.")
		bedrockAgentcoreControl_deleteGatewayCmd.MarkFlagRequired("gateway-identifier")
	})
	bedrockAgentcoreControlCmd.AddCommand(bedrockAgentcoreControl_deleteGatewayCmd)
}
