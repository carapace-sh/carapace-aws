package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_updateUserSecurityProfilesCmd = &cobra.Command{
	Use:   "update-user-security-profiles",
	Short: "Assigns the specified security profiles to the specified user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_updateUserSecurityProfilesCmd).Standalone()

	connect_updateUserSecurityProfilesCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_updateUserSecurityProfilesCmd.Flags().String("security-profile-ids", "", "The identifiers of the security profiles for the user.")
	connect_updateUserSecurityProfilesCmd.Flags().String("user-id", "", "The identifier of the user account.")
	connect_updateUserSecurityProfilesCmd.MarkFlagRequired("instance-id")
	connect_updateUserSecurityProfilesCmd.MarkFlagRequired("security-profile-ids")
	connect_updateUserSecurityProfilesCmd.MarkFlagRequired("user-id")
	connectCmd.AddCommand(connect_updateUserSecurityProfilesCmd)
}
