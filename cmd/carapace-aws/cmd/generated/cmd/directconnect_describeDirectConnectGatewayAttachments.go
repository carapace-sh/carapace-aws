package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var directconnect_describeDirectConnectGatewayAttachmentsCmd = &cobra.Command{
	Use:   "describe-direct-connect-gateway-attachments",
	Short: "Lists the attachments between your Direct Connect gateways and virtual interfaces.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(directconnect_describeDirectConnectGatewayAttachmentsCmd).Standalone()

	directconnect_describeDirectConnectGatewayAttachmentsCmd.Flags().String("direct-connect-gateway-id", "", "The ID of the Direct Connect gateway.")
	directconnect_describeDirectConnectGatewayAttachmentsCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
	directconnect_describeDirectConnectGatewayAttachmentsCmd.Flags().String("next-token", "", "The token provided in the previous call to retrieve the next page.")
	directconnect_describeDirectConnectGatewayAttachmentsCmd.Flags().String("virtual-interface-id", "", "The ID of the virtual interface.")
	directconnectCmd.AddCommand(directconnect_describeDirectConnectGatewayAttachmentsCmd)
}
