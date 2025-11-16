package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var omics_listReadSetActivationJobsCmd = &cobra.Command{
	Use:   "list-read-set-activation-jobs",
	Short: "Retrieves a list of read set activation jobs and returns the metadata in a JSON formatted output.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(omics_listReadSetActivationJobsCmd).Standalone()

	omics_listReadSetActivationJobsCmd.Flags().String("filter", "", "A filter to apply to the list.")
	omics_listReadSetActivationJobsCmd.Flags().String("max-results", "", "The maximum number of read set activation jobs to return in one page of results.")
	omics_listReadSetActivationJobsCmd.Flags().String("next-token", "", "Specify the pagination token from a previous request to retrieve the next page of results.")
	omics_listReadSetActivationJobsCmd.Flags().String("sequence-store-id", "", "The read set's sequence store ID.")
	omics_listReadSetActivationJobsCmd.MarkFlagRequired("sequence-store-id")
	omicsCmd.AddCommand(omics_listReadSetActivationJobsCmd)
}
