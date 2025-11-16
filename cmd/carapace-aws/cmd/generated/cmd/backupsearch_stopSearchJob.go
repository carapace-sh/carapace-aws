package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backupsearch_stopSearchJobCmd = &cobra.Command{
	Use:   "stop-search-job",
	Short: "This operations ends a search job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backupsearch_stopSearchJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(backupsearch_stopSearchJobCmd).Standalone()

		backupsearch_stopSearchJobCmd.Flags().String("search-job-identifier", "", "The unique string that specifies the search job.")
		backupsearch_stopSearchJobCmd.MarkFlagRequired("search-job-identifier")
	})
	backupsearchCmd.AddCommand(backupsearch_stopSearchJobCmd)
}
