package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backupsearch_listSearchJobsCmd = &cobra.Command{
	Use:   "list-search-jobs",
	Short: "This operation returns a list of search jobs belonging to an account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backupsearch_listSearchJobsCmd).Standalone()

	backupsearch_listSearchJobsCmd.Flags().String("by-status", "", "Include this parameter to filter list by search job status.")
	backupsearch_listSearchJobsCmd.Flags().String("max-results", "", "The maximum number of resource list items to be returned.")
	backupsearch_listSearchJobsCmd.Flags().String("next-token", "", "The next item following a partial list of returned search jobs.")
	backupsearchCmd.AddCommand(backupsearch_listSearchJobsCmd)
}
