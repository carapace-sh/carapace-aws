package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediaconnect_describeGatewayInstanceCmd = &cobra.Command{
	Use:   "describe-gateway-instance",
	Short: "Displays the details of an instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediaconnect_describeGatewayInstanceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediaconnect_describeGatewayInstanceCmd).Standalone()

		mediaconnect_describeGatewayInstanceCmd.Flags().String("gateway-instance-arn", "", "The Amazon Resource Name (ARN) of the gateway instance that you want to describe.")
		mediaconnect_describeGatewayInstanceCmd.MarkFlagRequired("gateway-instance-arn")
	})
	mediaconnectCmd.AddCommand(mediaconnect_describeGatewayInstanceCmd)
}
