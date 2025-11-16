package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var finspaceData_enableUserCmd = &cobra.Command{
	Use:   "enable-user",
	Short: "Allows the specified user to access the FinSpace web application and API.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(finspaceData_enableUserCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(finspaceData_enableUserCmd).Standalone()

		finspaceData_enableUserCmd.Flags().String("client-token", "", "A token that ensures idempotency.")
		finspaceData_enableUserCmd.Flags().String("user-id", "", "The unique identifier for the user that you want to activate.")
		finspaceData_enableUserCmd.MarkFlagRequired("user-id")
	})
	finspaceDataCmd.AddCommand(finspaceData_enableUserCmd)
}
