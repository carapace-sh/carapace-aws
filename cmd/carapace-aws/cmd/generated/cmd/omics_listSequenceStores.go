package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var omics_listSequenceStoresCmd = &cobra.Command{
	Use:   "list-sequence-stores",
	Short: "Retrieves a list of sequence stores and returns each sequence store's metadata.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(omics_listSequenceStoresCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(omics_listSequenceStoresCmd).Standalone()

		omics_listSequenceStoresCmd.Flags().String("filter", "", "A filter to apply to the list.")
		omics_listSequenceStoresCmd.Flags().String("max-results", "", "The maximum number of stores to return in one page of results.")
		omics_listSequenceStoresCmd.Flags().String("next-token", "", "Specify the pagination token from a previous request to retrieve the next page of results.")
	})
	omicsCmd.AddCommand(omics_listSequenceStoresCmd)
}
