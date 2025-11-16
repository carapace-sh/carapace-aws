package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var finspaceData_listUsersCmd = &cobra.Command{
	Use:   "list-users",
	Short: "Lists all available users in FinSpace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(finspaceData_listUsersCmd).Standalone()

	finspaceData_listUsersCmd.Flags().String("max-results", "", "The maximum number of results per page.")
	finspaceData_listUsersCmd.Flags().String("next-token", "", "A token that indicates where a results page should begin.")
	finspaceData_listUsersCmd.MarkFlagRequired("max-results")
	finspaceDataCmd.AddCommand(finspaceData_listUsersCmd)
}
