package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_describeGatewayCmd = &cobra.Command{
	Use:   "describe-gateway",
	Short: "Retrieves information about a gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_describeGatewayCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotsitewise_describeGatewayCmd).Standalone()

		iotsitewise_describeGatewayCmd.Flags().String("gateway-id", "", "The ID of the gateway device.")
		iotsitewise_describeGatewayCmd.MarkFlagRequired("gateway-id")
	})
	iotsitewiseCmd.AddCommand(iotsitewise_describeGatewayCmd)
}
