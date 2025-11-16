package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediaconnect_updateGatewayInstanceCmd = &cobra.Command{
	Use:   "update-gateway-instance",
	Short: "Updates an existing gateway instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediaconnect_updateGatewayInstanceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediaconnect_updateGatewayInstanceCmd).Standalone()

		mediaconnect_updateGatewayInstanceCmd.Flags().String("bridge-placement", "", "The state of the instance.")
		mediaconnect_updateGatewayInstanceCmd.Flags().String("gateway-instance-arn", "", "The Amazon Resource Name (ARN) of the gateway instance that you want to update.")
		mediaconnect_updateGatewayInstanceCmd.MarkFlagRequired("gateway-instance-arn")
	})
	mediaconnectCmd.AddCommand(mediaconnect_updateGatewayInstanceCmd)
}
