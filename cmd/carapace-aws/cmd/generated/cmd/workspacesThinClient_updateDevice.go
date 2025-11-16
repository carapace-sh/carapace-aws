package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesThinClient_updateDeviceCmd = &cobra.Command{
	Use:   "update-device",
	Short: "Updates a thin client device.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesThinClient_updateDeviceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workspacesThinClient_updateDeviceCmd).Standalone()

		workspacesThinClient_updateDeviceCmd.Flags().String("desired-software-set-id", "", "The ID of the software set to apply.")
		workspacesThinClient_updateDeviceCmd.Flags().String("id", "", "The ID of the device to update.")
		workspacesThinClient_updateDeviceCmd.Flags().String("name", "", "The name of the device to update.")
		workspacesThinClient_updateDeviceCmd.Flags().String("software-set-update-schedule", "", "An option to define if software updates should be applied within a maintenance window.")
		workspacesThinClient_updateDeviceCmd.MarkFlagRequired("id")
	})
	workspacesThinClientCmd.AddCommand(workspacesThinClient_updateDeviceCmd)
}
