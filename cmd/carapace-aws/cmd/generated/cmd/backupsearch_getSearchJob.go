package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backupsearch_getSearchJobCmd = &cobra.Command{
	Use:   "get-search-job",
	Short: "This operation retrieves metadata of a search job, including its progress.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backupsearch_getSearchJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(backupsearch_getSearchJobCmd).Standalone()

		backupsearch_getSearchJobCmd.Flags().String("search-job-identifier", "", "Required unique string that specifies the search job.")
		backupsearch_getSearchJobCmd.MarkFlagRequired("search-job-identifier")
	})
	backupsearchCmd.AddCommand(backupsearch_getSearchJobCmd)
}
