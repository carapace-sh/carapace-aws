package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediaconnect_updateBridgeOutputCmd = &cobra.Command{
	Use:   "update-bridge-output",
	Short: "Updates an existing bridge output.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediaconnect_updateBridgeOutputCmd).Standalone()

	mediaconnect_updateBridgeOutputCmd.Flags().String("bridge-arn", "", "The Amazon Resource Name (ARN) of the bridge that you want to update.")
	mediaconnect_updateBridgeOutputCmd.Flags().String("network-output", "", "The network of the bridge output.")
	mediaconnect_updateBridgeOutputCmd.Flags().String("output-name", "", "Tname of the output that you want to update.")
	mediaconnect_updateBridgeOutputCmd.MarkFlagRequired("bridge-arn")
	mediaconnect_updateBridgeOutputCmd.MarkFlagRequired("output-name")
	mediaconnectCmd.AddCommand(mediaconnect_updateBridgeOutputCmd)
}
