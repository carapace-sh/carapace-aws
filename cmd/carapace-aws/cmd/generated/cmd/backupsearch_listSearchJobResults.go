package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backupsearch_listSearchJobResultsCmd = &cobra.Command{
	Use:   "list-search-job-results",
	Short: "This operation returns a list of a specified search job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backupsearch_listSearchJobResultsCmd).Standalone()

	backupsearch_listSearchJobResultsCmd.Flags().String("max-results", "", "The maximum number of resource list items to be returned.")
	backupsearch_listSearchJobResultsCmd.Flags().String("next-token", "", "The next item following a partial list of returned search job results.")
	backupsearch_listSearchJobResultsCmd.Flags().String("search-job-identifier", "", "The unique string that specifies the search job.")
	backupsearch_listSearchJobResultsCmd.MarkFlagRequired("search-job-identifier")
	backupsearchCmd.AddCommand(backupsearch_listSearchJobResultsCmd)
}
