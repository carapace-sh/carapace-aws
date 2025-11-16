package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qbusiness_updateUserCmd = &cobra.Command{
	Use:   "update-user",
	Short: "Updates a information associated with a user id.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qbusiness_updateUserCmd).Standalone()

	qbusiness_updateUserCmd.Flags().String("application-id", "", "The identifier of the application the user is attached to.")
	qbusiness_updateUserCmd.Flags().String("user-aliases-to-delete", "", "The user aliases attached to the user id that are to be deleted.")
	qbusiness_updateUserCmd.Flags().String("user-aliases-to-update", "", "The user aliases attached to the user id that are to be updated.")
	qbusiness_updateUserCmd.Flags().String("user-id", "", "The email id attached to the user.")
	qbusiness_updateUserCmd.MarkFlagRequired("application-id")
	qbusiness_updateUserCmd.MarkFlagRequired("user-id")
	qbusinessCmd.AddCommand(qbusiness_updateUserCmd)
}
