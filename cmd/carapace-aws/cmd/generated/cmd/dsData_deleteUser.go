package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dsData_deleteUserCmd = &cobra.Command{
	Use:   "delete-user",
	Short: "Deletes a user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dsData_deleteUserCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dsData_deleteUserCmd).Standalone()

		dsData_deleteUserCmd.Flags().String("client-token", "", "A unique and case-sensitive identifier that you provide to make sure the idempotency of the request, so multiple identical calls have the same effect as one single call.")
		dsData_deleteUserCmd.Flags().String("directory-id", "", "The identifier (ID) of the directory that's associated with the user.")
		dsData_deleteUserCmd.Flags().String("samaccount-name", "", "The name of the user.")
		dsData_deleteUserCmd.MarkFlagRequired("directory-id")
		dsData_deleteUserCmd.MarkFlagRequired("samaccount-name")
	})
	dsDataCmd.AddCommand(dsData_deleteUserCmd)
}
