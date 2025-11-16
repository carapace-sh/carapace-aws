package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backupsearch_getSearchResultExportJobCmd = &cobra.Command{
	Use:   "get-search-result-export-job",
	Short: "This operation retrieves the metadata of an export job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backupsearch_getSearchResultExportJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(backupsearch_getSearchResultExportJobCmd).Standalone()

		backupsearch_getSearchResultExportJobCmd.Flags().String("export-job-identifier", "", "This is the unique string that identifies a specific export job.")
		backupsearch_getSearchResultExportJobCmd.MarkFlagRequired("export-job-identifier")
	})
	backupsearchCmd.AddCommand(backupsearch_getSearchResultExportJobCmd)
}
