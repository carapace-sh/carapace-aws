package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_adminGetDeviceCmd = &cobra.Command{
	Use:   "admin-get-device",
	Short: "Given the device key, returns details for a user's device.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_adminGetDeviceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cognitoIdp_adminGetDeviceCmd).Standalone()

		cognitoIdp_adminGetDeviceCmd.Flags().String("device-key", "", "The key of the device that you want to delete.")
		cognitoIdp_adminGetDeviceCmd.Flags().String("user-pool-id", "", "The ID of the user pool where the device owner is a user.")
		cognitoIdp_adminGetDeviceCmd.Flags().String("username", "", "The name of the user that you want to query or modify.")
		cognitoIdp_adminGetDeviceCmd.MarkFlagRequired("device-key")
		cognitoIdp_adminGetDeviceCmd.MarkFlagRequired("user-pool-id")
		cognitoIdp_adminGetDeviceCmd.MarkFlagRequired("username")
	})
	cognitoIdpCmd.AddCommand(cognitoIdp_adminGetDeviceCmd)
}
