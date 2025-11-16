package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dsData_updateUserCmd = &cobra.Command{
	Use:   "update-user",
	Short: "Updates user information.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dsData_updateUserCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dsData_updateUserCmd).Standalone()

		dsData_updateUserCmd.Flags().String("client-token", "", "A unique and case-sensitive identifier that you provide to make sure the idempotency of the request, so multiple identical calls have the same effect as one single call.")
		dsData_updateUserCmd.Flags().String("directory-id", "", "The identifier (ID) of the directory that's associated with the user.")
		dsData_updateUserCmd.Flags().String("email-address", "", "The email address of the user.")
		dsData_updateUserCmd.Flags().String("given-name", "", "The first name of the user.")
		dsData_updateUserCmd.Flags().String("other-attributes", "", "An expression that defines one or more attribute names with the data type and value of each attribute.")
		dsData_updateUserCmd.Flags().String("samaccount-name", "", "The name of the user.")
		dsData_updateUserCmd.Flags().String("surname", "", "The last name of the user.")
		dsData_updateUserCmd.Flags().String("update-type", "", "The type of update to be performed.")
		dsData_updateUserCmd.MarkFlagRequired("directory-id")
		dsData_updateUserCmd.MarkFlagRequired("samaccount-name")
	})
	dsDataCmd.AddCommand(dsData_updateUserCmd)
}
