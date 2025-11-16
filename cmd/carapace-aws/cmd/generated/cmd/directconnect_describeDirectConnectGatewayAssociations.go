package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var directconnect_describeDirectConnectGatewayAssociationsCmd = &cobra.Command{
	Use:   "describe-direct-connect-gateway-associations",
	Short: "Lists the associations between your Direct Connect gateways and virtual private gateways and transit gateways.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(directconnect_describeDirectConnectGatewayAssociationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(directconnect_describeDirectConnectGatewayAssociationsCmd).Standalone()

		directconnect_describeDirectConnectGatewayAssociationsCmd.Flags().String("associated-gateway-id", "", "The ID of the associated gateway.")
		directconnect_describeDirectConnectGatewayAssociationsCmd.Flags().String("association-id", "", "The ID of the Direct Connect gateway association.")
		directconnect_describeDirectConnectGatewayAssociationsCmd.Flags().String("direct-connect-gateway-id", "", "The ID of the Direct Connect gateway.")
		directconnect_describeDirectConnectGatewayAssociationsCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
		directconnect_describeDirectConnectGatewayAssociationsCmd.Flags().String("next-token", "", "The token provided in the previous call to retrieve the next page.")
		directconnect_describeDirectConnectGatewayAssociationsCmd.Flags().String("virtual-gateway-id", "", "The ID of the virtual private gateway or transit gateway.")
	})
	directconnectCmd.AddCommand(directconnect_describeDirectConnectGatewayAssociationsCmd)
}
