package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_describeGatewayInformationCmd = &cobra.Command{
	Use:   "describe-gateway-information",
	Short: "Returns metadata about a gateway such as its name, network interfaces, time zone, status, and software version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_describeGatewayInformationCmd).Standalone()

	storagegateway_describeGatewayInformationCmd.Flags().String("gateway-arn", "", "")
	storagegateway_describeGatewayInformationCmd.MarkFlagRequired("gateway-arn")
	storagegatewayCmd.AddCommand(storagegateway_describeGatewayInformationCmd)
}
