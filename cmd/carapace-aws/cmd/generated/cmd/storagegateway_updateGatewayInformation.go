package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_updateGatewayInformationCmd = &cobra.Command{
	Use:   "update-gateway-information",
	Short: "Updates a gateway's metadata, which includes the gateway's name, time zone, and metadata cache size.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_updateGatewayInformationCmd).Standalone()

	storagegateway_updateGatewayInformationCmd.Flags().String("cloud-watch-log-group-arn", "", "The Amazon Resource Name (ARN) of the Amazon CloudWatch log group that you want to use to monitor and log events in the gateway.")
	storagegateway_updateGatewayInformationCmd.Flags().String("gateway-arn", "", "")
	storagegateway_updateGatewayInformationCmd.Flags().String("gateway-capacity", "", "Specifies the size of the gateway's metadata cache.")
	storagegateway_updateGatewayInformationCmd.Flags().String("gateway-name", "", "")
	storagegateway_updateGatewayInformationCmd.Flags().String("gateway-timezone", "", "A value that indicates the time zone of the gateway.")
	storagegateway_updateGatewayInformationCmd.MarkFlagRequired("gateway-arn")
	storagegatewayCmd.AddCommand(storagegateway_updateGatewayInformationCmd)
}
