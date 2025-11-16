package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var directconnect_describeDirectConnectGatewaysCmd = &cobra.Command{
	Use:   "describe-direct-connect-gateways",
	Short: "Lists all your Direct Connect gateways or only the specified Direct Connect gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(directconnect_describeDirectConnectGatewaysCmd).Standalone()

	directconnect_describeDirectConnectGatewaysCmd.Flags().String("direct-connect-gateway-id", "", "The ID of the Direct Connect gateway.")
	directconnect_describeDirectConnectGatewaysCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
	directconnect_describeDirectConnectGatewaysCmd.Flags().String("next-token", "", "The token provided in the previous call to retrieve the next page.")
	directconnectCmd.AddCommand(directconnect_describeDirectConnectGatewaysCmd)
}
