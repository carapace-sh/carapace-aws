package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var braket_searchJobsCmd = &cobra.Command{
	Use:   "search-jobs",
	Short: "Searches for Amazon Braket hybrid jobs that match the specified filter values.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(braket_searchJobsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(braket_searchJobsCmd).Standalone()

		braket_searchJobsCmd.Flags().String("filters", "", "Array of SearchJobsFilter objects to use when searching for hybrid jobs.")
		braket_searchJobsCmd.Flags().String("max-results", "", "The maximum number of results to return in the response.")
		braket_searchJobsCmd.Flags().String("next-token", "", "A token used for pagination of results returned in the response.")
		braket_searchJobsCmd.MarkFlagRequired("filters")
	})
	braketCmd.AddCommand(braket_searchJobsCmd)
}
