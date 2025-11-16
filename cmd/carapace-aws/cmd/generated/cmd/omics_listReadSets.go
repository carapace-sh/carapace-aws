package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var omics_listReadSetsCmd = &cobra.Command{
	Use:   "list-read-sets",
	Short: "Retrieves a list of read sets from a sequence store ID and returns the metadata in JSON format.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(omics_listReadSetsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(omics_listReadSetsCmd).Standalone()

		omics_listReadSetsCmd.Flags().String("filter", "", "A filter to apply to the list.")
		omics_listReadSetsCmd.Flags().String("max-results", "", "The maximum number of read sets to return in one page of results.")
		omics_listReadSetsCmd.Flags().String("next-token", "", "Specify the pagination token from a previous request to retrieve the next page of results.")
		omics_listReadSetsCmd.Flags().String("sequence-store-id", "", "The jobs' sequence store ID.")
		omics_listReadSetsCmd.MarkFlagRequired("sequence-store-id")
	})
	omicsCmd.AddCommand(omics_listReadSetsCmd)
}
