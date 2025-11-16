package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediaconnect_createBridgeCmd = &cobra.Command{
	Use:   "create-bridge",
	Short: "Creates a new bridge.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediaconnect_createBridgeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediaconnect_createBridgeCmd).Standalone()

		mediaconnect_createBridgeCmd.Flags().String("egress-gateway-bridge", "", "An egress bridge is a cloud-to-ground bridge.")
		mediaconnect_createBridgeCmd.Flags().String("ingress-gateway-bridge", "", "An ingress bridge is a ground-to-cloud bridge.")
		mediaconnect_createBridgeCmd.Flags().String("name", "", "The name of the bridge.")
		mediaconnect_createBridgeCmd.Flags().String("outputs", "", "The outputs that you want to add to this bridge.")
		mediaconnect_createBridgeCmd.Flags().String("placement-arn", "", "The bridge placement Amazon Resource Number (ARN).")
		mediaconnect_createBridgeCmd.Flags().String("source-failover-config", "", "The settings for source failover.")
		mediaconnect_createBridgeCmd.Flags().String("sources", "", "The sources that you want to add to this bridge.")
		mediaconnect_createBridgeCmd.MarkFlagRequired("name")
		mediaconnect_createBridgeCmd.MarkFlagRequired("placement-arn")
		mediaconnect_createBridgeCmd.MarkFlagRequired("sources")
	})
	mediaconnectCmd.AddCommand(mediaconnect_createBridgeCmd)
}
