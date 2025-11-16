package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentcoreControl_getGatewayCmd = &cobra.Command{
	Use:   "get-gateway",
	Short: "Retrieves information about a specific Gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentcoreControl_getGatewayCmd).Standalone()

	bedrockAgentcoreControl_getGatewayCmd.Flags().String("gateway-identifier", "", "The identifier of the gateway to retrieve.")
	bedrockAgentcoreControl_getGatewayCmd.MarkFlagRequired("gateway-identifier")
	bedrockAgentcoreControlCmd.AddCommand(bedrockAgentcoreControl_getGatewayCmd)
}
