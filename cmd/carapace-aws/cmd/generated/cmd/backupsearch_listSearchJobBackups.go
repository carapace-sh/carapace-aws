package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backupsearch_listSearchJobBackupsCmd = &cobra.Command{
	Use:   "list-search-job-backups",
	Short: "This operation returns a list of all backups (recovery points) in a paginated format that were included in the search job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backupsearch_listSearchJobBackupsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(backupsearch_listSearchJobBackupsCmd).Standalone()

		backupsearch_listSearchJobBackupsCmd.Flags().String("max-results", "", "The maximum number of resource list items to be returned.")
		backupsearch_listSearchJobBackupsCmd.Flags().String("next-token", "", "The next item following a partial list of returned backups included in a search job.")
		backupsearch_listSearchJobBackupsCmd.Flags().String("search-job-identifier", "", "The unique string that specifies the search job.")
		backupsearch_listSearchJobBackupsCmd.MarkFlagRequired("search-job-identifier")
	})
	backupsearchCmd.AddCommand(backupsearch_listSearchJobBackupsCmd)
}
