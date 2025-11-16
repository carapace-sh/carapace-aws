package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_deleteUserCmd = &cobra.Command{
	Use:   "delete-user",
	Short: "Deletes a user account from the specified Amazon Connect instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_deleteUserCmd).Standalone()

	connect_deleteUserCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_deleteUserCmd.Flags().String("user-id", "", "The identifier of the user.")
	connect_deleteUserCmd.MarkFlagRequired("instance-id")
	connect_deleteUserCmd.MarkFlagRequired("user-id")
	connectCmd.AddCommand(connect_deleteUserCmd)
}
