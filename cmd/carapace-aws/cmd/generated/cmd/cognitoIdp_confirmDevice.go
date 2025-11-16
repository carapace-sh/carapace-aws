package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_confirmDeviceCmd = &cobra.Command{
	Use:   "confirm-device",
	Short: "Confirms a device that a user wants to remember.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_confirmDeviceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cognitoIdp_confirmDeviceCmd).Standalone()

		cognitoIdp_confirmDeviceCmd.Flags().String("access-token", "", "A valid access token that Amazon Cognito issued to the currently signed-in user.")
		cognitoIdp_confirmDeviceCmd.Flags().String("device-key", "", "The unique identifier, or device key, of the device that you want to update the status for.")
		cognitoIdp_confirmDeviceCmd.Flags().String("device-name", "", "A friendly name for the device, for example `MyMobilePhone`.")
		cognitoIdp_confirmDeviceCmd.Flags().String("device-secret-verifier-config", "", "The configuration of the device secret verifier.")
		cognitoIdp_confirmDeviceCmd.MarkFlagRequired("access-token")
		cognitoIdp_confirmDeviceCmd.MarkFlagRequired("device-key")
	})
	cognitoIdpCmd.AddCommand(cognitoIdp_confirmDeviceCmd)
}
