package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_forgetDeviceCmd = &cobra.Command{
	Use:   "forget-device",
	Short: "Given a device key, deletes a remembered device as the currently signed-in user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_forgetDeviceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cognitoIdp_forgetDeviceCmd).Standalone()

		cognitoIdp_forgetDeviceCmd.Flags().String("access-token", "", "A valid access token that Amazon Cognito issued to the currently signed-in user.")
		cognitoIdp_forgetDeviceCmd.Flags().String("device-key", "", "The unique identifier, or device key, of the device that the user wants to forget.")
		cognitoIdp_forgetDeviceCmd.MarkFlagRequired("device-key")
	})
	cognitoIdpCmd.AddCommand(cognitoIdp_forgetDeviceCmd)
}
