package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_listBulkImportJobsCmd = &cobra.Command{
	Use:   "list-bulk-import-jobs",
	Short: "Retrieves a paginated list of bulk import job requests.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_listBulkImportJobsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotsitewise_listBulkImportJobsCmd).Standalone()

		iotsitewise_listBulkImportJobsCmd.Flags().String("filter", "", "You can use a filter to select the bulk import jobs that you want to retrieve.")
		iotsitewise_listBulkImportJobsCmd.Flags().String("max-results", "", "The maximum number of results to return for each paginated request.")
		iotsitewise_listBulkImportJobsCmd.Flags().String("next-token", "", "The token to be used for the next set of paginated results.")
	})
	iotsitewiseCmd.AddCommand(iotsitewise_listBulkImportJobsCmd)
}
