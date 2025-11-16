package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_updateUserRoutingProfileCmd = &cobra.Command{
	Use:   "update-user-routing-profile",
	Short: "Assigns the specified routing profile to the specified user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_updateUserRoutingProfileCmd).Standalone()

	connect_updateUserRoutingProfileCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_updateUserRoutingProfileCmd.Flags().String("routing-profile-id", "", "The identifier of the routing profile for the user.")
	connect_updateUserRoutingProfileCmd.Flags().String("user-id", "", "The identifier of the user account.")
	connect_updateUserRoutingProfileCmd.MarkFlagRequired("instance-id")
	connect_updateUserRoutingProfileCmd.MarkFlagRequired("routing-profile-id")
	connect_updateUserRoutingProfileCmd.MarkFlagRequired("user-id")
	connectCmd.AddCommand(connect_updateUserRoutingProfileCmd)
}
