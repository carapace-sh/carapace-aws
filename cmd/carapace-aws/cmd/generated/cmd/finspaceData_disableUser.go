package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var finspaceData_disableUserCmd = &cobra.Command{
	Use:   "disable-user",
	Short: "Denies access to the FinSpace web application and API for the specified user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(finspaceData_disableUserCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(finspaceData_disableUserCmd).Standalone()

		finspaceData_disableUserCmd.Flags().String("client-token", "", "A token that ensures idempotency.")
		finspaceData_disableUserCmd.Flags().String("user-id", "", "The unique identifier for the user that you want to deactivate.")
		finspaceData_disableUserCmd.MarkFlagRequired("user-id")
	})
	finspaceDataCmd.AddCommand(finspaceData_disableUserCmd)
}
