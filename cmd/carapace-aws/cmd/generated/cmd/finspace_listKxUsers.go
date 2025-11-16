package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var finspace_listKxUsersCmd = &cobra.Command{
	Use:   "list-kx-users",
	Short: "Lists all the users in a kdb environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(finspace_listKxUsersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(finspace_listKxUsersCmd).Standalone()

		finspace_listKxUsersCmd.Flags().String("environment-id", "", "A unique identifier for the kdb environment.")
		finspace_listKxUsersCmd.Flags().String("max-results", "", "The maximum number of results to return in this request.")
		finspace_listKxUsersCmd.Flags().String("next-token", "", "A token that indicates where a results page should begin.")
		finspace_listKxUsersCmd.MarkFlagRequired("environment-id")
	})
	finspaceCmd.AddCommand(finspace_listKxUsersCmd)
}
