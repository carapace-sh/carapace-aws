package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanrooms_listProtectedJobsCmd = &cobra.Command{
	Use:   "list-protected-jobs",
	Short: "Lists protected jobs, sorted by most recent job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanrooms_listProtectedJobsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cleanrooms_listProtectedJobsCmd).Standalone()

		cleanrooms_listProtectedJobsCmd.Flags().String("max-results", "", "The maximum number of results that are returned for an API request call.")
		cleanrooms_listProtectedJobsCmd.Flags().String("membership-identifier", "", "The identifier for the membership in the collaboration.")
		cleanrooms_listProtectedJobsCmd.Flags().String("next-token", "", "The pagination token that's used to fetch the next set of results.")
		cleanrooms_listProtectedJobsCmd.Flags().String("status", "", "A filter on the status of the protected job.")
		cleanrooms_listProtectedJobsCmd.MarkFlagRequired("membership-identifier")
	})
	cleanroomsCmd.AddCommand(cleanrooms_listProtectedJobsCmd)
}
