package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dsData_createUserCmd = &cobra.Command{
	Use:   "create-user",
	Short: "Creates a new user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dsData_createUserCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dsData_createUserCmd).Standalone()

		dsData_createUserCmd.Flags().String("client-token", "", "A unique and case-sensitive identifier that you provide to make sure the idempotency of the request, so multiple identical calls have the same effect as one single call.")
		dsData_createUserCmd.Flags().String("directory-id", "", "The identifier (ID) of the directory that’s associated with the user.")
		dsData_createUserCmd.Flags().String("email-address", "", "The email address of the user.")
		dsData_createUserCmd.Flags().String("given-name", "", "The first name of the user.")
		dsData_createUserCmd.Flags().String("other-attributes", "", "An expression that defines one or more attribute names with the data type and value of each attribute.")
		dsData_createUserCmd.Flags().String("samaccount-name", "", "The name of the user.")
		dsData_createUserCmd.Flags().String("surname", "", "The last name of the user.")
		dsData_createUserCmd.MarkFlagRequired("directory-id")
		dsData_createUserCmd.MarkFlagRequired("samaccount-name")
	})
	dsDataCmd.AddCommand(dsData_createUserCmd)
}
