package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var omics_listReferencesCmd = &cobra.Command{
	Use:   "list-references",
	Short: "Retrieves the metadata of one or more reference genomes in a reference store.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(omics_listReferencesCmd).Standalone()

	omics_listReferencesCmd.Flags().String("filter", "", "A filter to apply to the list.")
	omics_listReferencesCmd.Flags().String("max-results", "", "The maximum number of references to return in one page of results.")
	omics_listReferencesCmd.Flags().String("next-token", "", "Specify the pagination token from a previous request to retrieve the next page of results.")
	omics_listReferencesCmd.Flags().String("reference-store-id", "", "The references' reference store ID.")
	omics_listReferencesCmd.MarkFlagRequired("reference-store-id")
	omicsCmd.AddCommand(omics_listReferencesCmd)
}
