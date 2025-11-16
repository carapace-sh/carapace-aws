package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_getDeviceCmd = &cobra.Command{
	Use:   "get-device",
	Short: "Given a device key, returns information about a remembered device for the current user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_getDeviceCmd).Standalone()

	cognitoIdp_getDeviceCmd.Flags().String("access-token", "", "A valid access token that Amazon Cognito issued to the currently signed-in user.")
	cognitoIdp_getDeviceCmd.Flags().String("device-key", "", "The key of the device that you want to get information about.")
	cognitoIdp_getDeviceCmd.MarkFlagRequired("device-key")
	cognitoIdpCmd.AddCommand(cognitoIdp_getDeviceCmd)
}
