package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesThinClient_deregisterDeviceCmd = &cobra.Command{
	Use:   "deregister-device",
	Short: "Deregisters a thin client device.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesThinClient_deregisterDeviceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workspacesThinClient_deregisterDeviceCmd).Standalone()

		workspacesThinClient_deregisterDeviceCmd.Flags().String("client-token", "", "Specifies a unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		workspacesThinClient_deregisterDeviceCmd.Flags().String("id", "", "The ID of the device to deregister.")
		workspacesThinClient_deregisterDeviceCmd.Flags().String("target-device-status", "", "The desired new status for the device.")
		workspacesThinClient_deregisterDeviceCmd.MarkFlagRequired("id")
	})
	workspacesThinClientCmd.AddCommand(workspacesThinClient_deregisterDeviceCmd)
}
