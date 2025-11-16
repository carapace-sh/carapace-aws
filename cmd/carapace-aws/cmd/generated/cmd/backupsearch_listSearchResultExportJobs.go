package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backupsearch_listSearchResultExportJobsCmd = &cobra.Command{
	Use:   "list-search-result-export-jobs",
	Short: "This operation exports search results of a search job to a specified destination S3 bucket.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backupsearch_listSearchResultExportJobsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(backupsearch_listSearchResultExportJobsCmd).Standalone()

		backupsearch_listSearchResultExportJobsCmd.Flags().String("max-results", "", "The maximum number of resource list items to be returned.")
		backupsearch_listSearchResultExportJobsCmd.Flags().String("next-token", "", "The next item following a partial list of returned backups included in a search job.")
		backupsearch_listSearchResultExportJobsCmd.Flags().String("search-job-identifier", "", "The unique string that specifies the search job.")
		backupsearch_listSearchResultExportJobsCmd.Flags().String("status", "", "The search jobs to be included in the export job can be filtered by including this parameter.")
	})
	backupsearchCmd.AddCommand(backupsearch_listSearchResultExportJobsCmd)
}
