package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediaconnect_describeGatewayCmd = &cobra.Command{
	Use:   "describe-gateway",
	Short: "Displays the details of a gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediaconnect_describeGatewayCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediaconnect_describeGatewayCmd).Standalone()

		mediaconnect_describeGatewayCmd.Flags().String("gateway-arn", "", "The ARN of the gateway that you want to describe.")
		mediaconnect_describeGatewayCmd.MarkFlagRequired("gateway-arn")
	})
	mediaconnectCmd.AddCommand(mediaconnect_describeGatewayCmd)
}
