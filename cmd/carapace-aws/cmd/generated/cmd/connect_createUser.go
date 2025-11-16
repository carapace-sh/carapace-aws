package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_createUserCmd = &cobra.Command{
	Use:   "create-user",
	Short: "Creates a user account for the specified Amazon Connect instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_createUserCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_createUserCmd).Standalone()

		connect_createUserCmd.Flags().String("directory-user-id", "", "The identifier of the user account in the directory used for identity management.")
		connect_createUserCmd.Flags().String("hierarchy-group-id", "", "The identifier of the hierarchy group for the user.")
		connect_createUserCmd.Flags().String("identity-info", "", "The information about the identity of the user.")
		connect_createUserCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_createUserCmd.Flags().String("password", "", "The password for the user account.")
		connect_createUserCmd.Flags().String("phone-config", "", "The phone settings for the user.")
		connect_createUserCmd.Flags().String("routing-profile-id", "", "The identifier of the routing profile for the user.")
		connect_createUserCmd.Flags().String("security-profile-ids", "", "The identifier of the security profile for the user.")
		connect_createUserCmd.Flags().String("tags", "", "The tags used to organize, track, or control access for this resource.")
		connect_createUserCmd.Flags().String("username", "", "The user name for the account.")
		connect_createUserCmd.MarkFlagRequired("instance-id")
		connect_createUserCmd.MarkFlagRequired("phone-config")
		connect_createUserCmd.MarkFlagRequired("routing-profile-id")
		connect_createUserCmd.MarkFlagRequired("security-profile-ids")
		connect_createUserCmd.MarkFlagRequired("username")
	})
	connectCmd.AddCommand(connect_createUserCmd)
}
