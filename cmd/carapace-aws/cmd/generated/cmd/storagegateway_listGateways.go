package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_listGatewaysCmd = &cobra.Command{
	Use:   "list-gateways",
	Short: "Lists gateways owned by an Amazon Web Services account in an Amazon Web Services Region specified in the request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_listGatewaysCmd).Standalone()

	storagegateway_listGatewaysCmd.Flags().String("limit", "", "Specifies that the list of gateways returned be limited to the specified number of items.")
	storagegateway_listGatewaysCmd.Flags().String("marker", "", "An opaque string that indicates the position at which to begin the returned list of gateways.")
	storagegatewayCmd.AddCommand(storagegateway_listGatewaysCmd)
}
