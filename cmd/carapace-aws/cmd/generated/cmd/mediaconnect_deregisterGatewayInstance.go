package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediaconnect_deregisterGatewayInstanceCmd = &cobra.Command{
	Use:   "deregister-gateway-instance",
	Short: "Deregisters an instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediaconnect_deregisterGatewayInstanceCmd).Standalone()

	mediaconnect_deregisterGatewayInstanceCmd.Flags().Bool("force", false, "Force the deregistration of an instance.")
	mediaconnect_deregisterGatewayInstanceCmd.Flags().String("gateway-instance-arn", "", "The Amazon Resource Name (ARN) of the gateway that contains the instance that you want to deregister.")
	mediaconnect_deregisterGatewayInstanceCmd.Flags().Bool("no-force", false, "Force the deregistration of an instance.")
	mediaconnect_deregisterGatewayInstanceCmd.MarkFlagRequired("gateway-instance-arn")
	mediaconnect_deregisterGatewayInstanceCmd.Flag("no-force").Hidden = true
	mediaconnectCmd.AddCommand(mediaconnect_deregisterGatewayInstanceCmd)
}
