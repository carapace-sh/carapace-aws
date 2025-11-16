package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediaconnect_updateBridgeSourceCmd = &cobra.Command{
	Use:   "update-bridge-source",
	Short: "Updates an existing bridge source.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediaconnect_updateBridgeSourceCmd).Standalone()

	mediaconnect_updateBridgeSourceCmd.Flags().String("bridge-arn", "", "The Amazon Resource Name (ARN) of the bridge that you want to update.")
	mediaconnect_updateBridgeSourceCmd.Flags().String("flow-source", "", "The name of the flow that you want to update.")
	mediaconnect_updateBridgeSourceCmd.Flags().String("network-source", "", "The network for the bridge source.")
	mediaconnect_updateBridgeSourceCmd.Flags().String("source-name", "", "The name of the source that you want to update.")
	mediaconnect_updateBridgeSourceCmd.MarkFlagRequired("bridge-arn")
	mediaconnect_updateBridgeSourceCmd.MarkFlagRequired("source-name")
	mediaconnectCmd.AddCommand(mediaconnect_updateBridgeSourceCmd)
}
