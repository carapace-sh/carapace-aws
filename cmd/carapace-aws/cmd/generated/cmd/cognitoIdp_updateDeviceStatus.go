package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_updateDeviceStatusCmd = &cobra.Command{
	Use:   "update-device-status",
	Short: "Updates the status of a the currently signed-in user's device so that it is marked as remembered or not remembered for the purpose of device authentication.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_updateDeviceStatusCmd).Standalone()

	cognitoIdp_updateDeviceStatusCmd.Flags().String("access-token", "", "A valid access token that Amazon Cognito issued to the currently signed-in user.")
	cognitoIdp_updateDeviceStatusCmd.Flags().String("device-key", "", "The device key of the device you want to update, for example `us-west-2_a1b2c3d4-5678-90ab-cdef-EXAMPLE11111`.")
	cognitoIdp_updateDeviceStatusCmd.Flags().String("device-remembered-status", "", "To enable device authentication with the specified device, set to `remembered`.To disable, set to `not_remembered`.")
	cognitoIdp_updateDeviceStatusCmd.MarkFlagRequired("access-token")
	cognitoIdp_updateDeviceStatusCmd.MarkFlagRequired("device-key")
	cognitoIdpCmd.AddCommand(cognitoIdp_updateDeviceStatusCmd)
}
