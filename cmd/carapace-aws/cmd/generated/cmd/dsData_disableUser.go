package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dsData_disableUserCmd = &cobra.Command{
	Use:   "disable-user",
	Short: "Deactivates an active user account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dsData_disableUserCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dsData_disableUserCmd).Standalone()

		dsData_disableUserCmd.Flags().String("client-token", "", "A unique and case-sensitive identifier that you provide to make sure the idempotency of the request, so multiple identical calls have the same effect as one single call.")
		dsData_disableUserCmd.Flags().String("directory-id", "", "The identifier (ID) of the directory that's associated with the user.")
		dsData_disableUserCmd.Flags().String("samaccount-name", "", "The name of the user.")
		dsData_disableUserCmd.MarkFlagRequired("directory-id")
		dsData_disableUserCmd.MarkFlagRequired("samaccount-name")
	})
	dsDataCmd.AddCommand(dsData_disableUserCmd)
}
