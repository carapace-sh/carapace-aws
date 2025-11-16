package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_updateUserPhoneConfigCmd = &cobra.Command{
	Use:   "update-user-phone-config",
	Short: "Updates the phone configuration settings for the specified user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_updateUserPhoneConfigCmd).Standalone()

	connect_updateUserPhoneConfigCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_updateUserPhoneConfigCmd.Flags().String("phone-config", "", "Information about phone configuration settings for the user.")
	connect_updateUserPhoneConfigCmd.Flags().String("user-id", "", "The identifier of the user account.")
	connect_updateUserPhoneConfigCmd.MarkFlagRequired("instance-id")
	connect_updateUserPhoneConfigCmd.MarkFlagRequired("phone-config")
	connect_updateUserPhoneConfigCmd.MarkFlagRequired("user-id")
	connectCmd.AddCommand(connect_updateUserPhoneConfigCmd)
}
