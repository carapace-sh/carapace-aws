package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backupsearch_startSearchResultExportJobCmd = &cobra.Command{
	Use:   "start-search-result-export-job",
	Short: "This operations starts a job to export the results of search job to a designated S3 bucket.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backupsearch_startSearchResultExportJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(backupsearch_startSearchResultExportJobCmd).Standalone()

		backupsearch_startSearchResultExportJobCmd.Flags().String("client-token", "", "Include this parameter to allow multiple identical calls for idempotency.")
		backupsearch_startSearchResultExportJobCmd.Flags().String("export-specification", "", "This specification contains a required string of the destination bucket; optionally, you can include the destination prefix.")
		backupsearch_startSearchResultExportJobCmd.Flags().String("role-arn", "", "This parameter specifies the role ARN used to start the search results export jobs.")
		backupsearch_startSearchResultExportJobCmd.Flags().String("search-job-identifier", "", "The unique string that specifies the search job.")
		backupsearch_startSearchResultExportJobCmd.Flags().String("tags", "", "Optional tags to include.")
		backupsearch_startSearchResultExportJobCmd.MarkFlagRequired("export-specification")
		backupsearch_startSearchResultExportJobCmd.MarkFlagRequired("search-job-identifier")
	})
	backupsearchCmd.AddCommand(backupsearch_startSearchResultExportJobCmd)
}
