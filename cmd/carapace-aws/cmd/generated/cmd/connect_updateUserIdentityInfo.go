package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_updateUserIdentityInfoCmd = &cobra.Command{
	Use:   "update-user-identity-info",
	Short: "Updates the identity information for the specified user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_updateUserIdentityInfoCmd).Standalone()

	connect_updateUserIdentityInfoCmd.Flags().String("identity-info", "", "The identity information for the user.")
	connect_updateUserIdentityInfoCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_updateUserIdentityInfoCmd.Flags().String("user-id", "", "The identifier of the user account.")
	connect_updateUserIdentityInfoCmd.MarkFlagRequired("identity-info")
	connect_updateUserIdentityInfoCmd.MarkFlagRequired("instance-id")
	connect_updateUserIdentityInfoCmd.MarkFlagRequired("user-id")
	connectCmd.AddCommand(connect_updateUserIdentityInfoCmd)
}
