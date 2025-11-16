package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var finspaceData_getUserCmd = &cobra.Command{
	Use:   "get-user",
	Short: "Retrieves details for a specific user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(finspaceData_getUserCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(finspaceData_getUserCmd).Standalone()

		finspaceData_getUserCmd.Flags().String("user-id", "", "The unique identifier of the user to get data for.")
		finspaceData_getUserCmd.MarkFlagRequired("user-id")
	})
	finspaceDataCmd.AddCommand(finspaceData_getUserCmd)
}
