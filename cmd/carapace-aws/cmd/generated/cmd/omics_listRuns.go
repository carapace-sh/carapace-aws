package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var omics_listRunsCmd = &cobra.Command{
	Use:   "list-runs",
	Short: "Retrieves a list of runs and returns each run's metadata and status.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(omics_listRunsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(omics_listRunsCmd).Standalone()

		omics_listRunsCmd.Flags().String("max-results", "", "The maximum number of runs to return in one page of results.")
		omics_listRunsCmd.Flags().String("name", "", "Filter the list by run name.")
		omics_listRunsCmd.Flags().String("run-group-id", "", "Filter the list by run group ID.")
		omics_listRunsCmd.Flags().String("starting-token", "", "Specify the pagination token from a previous request to retrieve the next page of results.")
		omics_listRunsCmd.Flags().String("status", "", "The status of a run.")
	})
	omicsCmd.AddCommand(omics_listRunsCmd)
}
