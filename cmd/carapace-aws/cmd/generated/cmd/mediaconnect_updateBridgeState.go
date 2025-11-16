package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediaconnect_updateBridgeStateCmd = &cobra.Command{
	Use:   "update-bridge-state",
	Short: "Updates the bridge state.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediaconnect_updateBridgeStateCmd).Standalone()

	mediaconnect_updateBridgeStateCmd.Flags().String("bridge-arn", "", "The Amazon Resource Name (ARN) of the bridge that you want to update the state of.")
	mediaconnect_updateBridgeStateCmd.Flags().String("desired-state", "", "The desired state for the bridge.")
	mediaconnect_updateBridgeStateCmd.MarkFlagRequired("bridge-arn")
	mediaconnect_updateBridgeStateCmd.MarkFlagRequired("desired-state")
	mediaconnectCmd.AddCommand(mediaconnect_updateBridgeStateCmd)
}
