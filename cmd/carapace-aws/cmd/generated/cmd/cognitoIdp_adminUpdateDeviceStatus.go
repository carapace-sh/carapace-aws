package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_adminUpdateDeviceStatusCmd = &cobra.Command{
	Use:   "admin-update-device-status",
	Short: "Updates the status of a user's device so that it is marked as remembered or not remembered for the purpose of device authentication.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_adminUpdateDeviceStatusCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cognitoIdp_adminUpdateDeviceStatusCmd).Standalone()

		cognitoIdp_adminUpdateDeviceStatusCmd.Flags().String("device-key", "", "The unique identifier, or device key, of the device that you want to update the status for.")
		cognitoIdp_adminUpdateDeviceStatusCmd.Flags().String("device-remembered-status", "", "To enable device authentication with the specified device, set to `remembered`.To disable, set to `not_remembered`.")
		cognitoIdp_adminUpdateDeviceStatusCmd.Flags().String("user-pool-id", "", "The ID of the user pool where you want to change a user's device status.")
		cognitoIdp_adminUpdateDeviceStatusCmd.Flags().String("username", "", "The name of the user that you want to query or modify.")
		cognitoIdp_adminUpdateDeviceStatusCmd.MarkFlagRequired("device-key")
		cognitoIdp_adminUpdateDeviceStatusCmd.MarkFlagRequired("user-pool-id")
		cognitoIdp_adminUpdateDeviceStatusCmd.MarkFlagRequired("username")
	})
	cognitoIdpCmd.AddCommand(cognitoIdp_adminUpdateDeviceStatusCmd)
}
