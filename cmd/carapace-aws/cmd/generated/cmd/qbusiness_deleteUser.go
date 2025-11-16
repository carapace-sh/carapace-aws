package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qbusiness_deleteUserCmd = &cobra.Command{
	Use:   "delete-user",
	Short: "Deletes a user by email id.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qbusiness_deleteUserCmd).Standalone()

	qbusiness_deleteUserCmd.Flags().String("application-id", "", "The identifier of the application from which the user is being deleted.")
	qbusiness_deleteUserCmd.Flags().String("user-id", "", "The user email being deleted.")
	qbusiness_deleteUserCmd.MarkFlagRequired("application-id")
	qbusiness_deleteUserCmd.MarkFlagRequired("user-id")
	qbusinessCmd.AddCommand(qbusiness_deleteUserCmd)
}
