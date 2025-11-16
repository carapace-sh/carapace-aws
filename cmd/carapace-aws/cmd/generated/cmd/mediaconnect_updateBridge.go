package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediaconnect_updateBridgeCmd = &cobra.Command{
	Use:   "update-bridge",
	Short: "Updates the bridge.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediaconnect_updateBridgeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediaconnect_updateBridgeCmd).Standalone()

		mediaconnect_updateBridgeCmd.Flags().String("bridge-arn", "", "TheAmazon Resource Name (ARN) of the bridge that you want to update.")
		mediaconnect_updateBridgeCmd.Flags().String("egress-gateway-bridge", "", "A cloud-to-ground bridge.")
		mediaconnect_updateBridgeCmd.Flags().String("ingress-gateway-bridge", "", "A ground-to-cloud bridge.")
		mediaconnect_updateBridgeCmd.Flags().String("source-failover-config", "", "The settings for source failover.")
		mediaconnect_updateBridgeCmd.MarkFlagRequired("bridge-arn")
	})
	mediaconnectCmd.AddCommand(mediaconnect_updateBridgeCmd)
}
