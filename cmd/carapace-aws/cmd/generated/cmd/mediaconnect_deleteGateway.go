package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediaconnect_deleteGatewayCmd = &cobra.Command{
	Use:   "delete-gateway",
	Short: "Deletes a gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediaconnect_deleteGatewayCmd).Standalone()

	mediaconnect_deleteGatewayCmd.Flags().String("gateway-arn", "", "The Amazon Resource Name (ARN) of the gateway that you want to delete.")
	mediaconnect_deleteGatewayCmd.MarkFlagRequired("gateway-arn")
	mediaconnectCmd.AddCommand(mediaconnect_deleteGatewayCmd)
}
