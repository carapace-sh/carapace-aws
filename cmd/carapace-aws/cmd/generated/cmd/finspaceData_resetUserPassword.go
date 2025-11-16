package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var finspaceData_resetUserPasswordCmd = &cobra.Command{
	Use:   "reset-user-password",
	Short: "Resets the password for a specified user ID and generates a temporary one.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(finspaceData_resetUserPasswordCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(finspaceData_resetUserPasswordCmd).Standalone()

		finspaceData_resetUserPasswordCmd.Flags().String("client-token", "", "A token that ensures idempotency.")
		finspaceData_resetUserPasswordCmd.Flags().String("user-id", "", "The unique identifier of the user that a temporary password is requested for.")
		finspaceData_resetUserPasswordCmd.MarkFlagRequired("user-id")
	})
	finspaceDataCmd.AddCommand(finspaceData_resetUserPasswordCmd)
}
