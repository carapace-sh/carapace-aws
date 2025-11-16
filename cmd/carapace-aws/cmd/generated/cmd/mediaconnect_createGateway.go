package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediaconnect_createGatewayCmd = &cobra.Command{
	Use:   "create-gateway",
	Short: "Creates a new gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediaconnect_createGatewayCmd).Standalone()

	mediaconnect_createGatewayCmd.Flags().String("egress-cidr-blocks", "", "The range of IP addresses that are allowed to contribute content or initiate output requests for flows communicating with this gateway.")
	mediaconnect_createGatewayCmd.Flags().String("name", "", "The name of the gateway.")
	mediaconnect_createGatewayCmd.Flags().String("networks", "", "The list of networks that you want to add to the gateway.")
	mediaconnect_createGatewayCmd.MarkFlagRequired("egress-cidr-blocks")
	mediaconnect_createGatewayCmd.MarkFlagRequired("name")
	mediaconnect_createGatewayCmd.MarkFlagRequired("networks")
	mediaconnectCmd.AddCommand(mediaconnect_createGatewayCmd)
}
