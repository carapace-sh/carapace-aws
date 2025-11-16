package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_adminForgetDeviceCmd = &cobra.Command{
	Use:   "admin-forget-device",
	Short: "Forgets, or deletes, a remembered device from a user's profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_adminForgetDeviceCmd).Standalone()

	cognitoIdp_adminForgetDeviceCmd.Flags().String("device-key", "", "The key ID of the device that you want to delete.")
	cognitoIdp_adminForgetDeviceCmd.Flags().String("user-pool-id", "", "The ID of the user pool where the device owner is a user.")
	cognitoIdp_adminForgetDeviceCmd.Flags().String("username", "", "The name of the user that you want to query or modify.")
	cognitoIdp_adminForgetDeviceCmd.MarkFlagRequired("device-key")
	cognitoIdp_adminForgetDeviceCmd.MarkFlagRequired("user-pool-id")
	cognitoIdp_adminForgetDeviceCmd.MarkFlagRequired("username")
	cognitoIdpCmd.AddCommand(cognitoIdp_adminForgetDeviceCmd)
}
