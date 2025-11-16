package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesThinClient_deleteDeviceCmd = &cobra.Command{
	Use:   "delete-device",
	Short: "Deletes a thin client device.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesThinClient_deleteDeviceCmd).Standalone()

	workspacesThinClient_deleteDeviceCmd.Flags().String("client-token", "", "Specifies a unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	workspacesThinClient_deleteDeviceCmd.Flags().String("id", "", "The ID of the device to delete.")
	workspacesThinClient_deleteDeviceCmd.MarkFlagRequired("id")
	workspacesThinClientCmd.AddCommand(workspacesThinClient_deleteDeviceCmd)
}
