package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var omics_listReferenceStoresCmd = &cobra.Command{
	Use:   "list-reference-stores",
	Short: "Retrieves a list of reference stores linked to your account and returns their metadata in JSON format.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(omics_listReferenceStoresCmd).Standalone()

	omics_listReferenceStoresCmd.Flags().String("filter", "", "A filter to apply to the list.")
	omics_listReferenceStoresCmd.Flags().String("max-results", "", "The maximum number of stores to return in one page of results.")
	omics_listReferenceStoresCmd.Flags().String("next-token", "", "Specify the pagination token from a previous request to retrieve the next page of results.")
	omicsCmd.AddCommand(omics_listReferenceStoresCmd)
}
