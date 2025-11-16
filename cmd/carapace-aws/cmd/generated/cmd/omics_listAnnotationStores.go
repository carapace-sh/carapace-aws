package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var omics_listAnnotationStoresCmd = &cobra.Command{
	Use:   "list-annotation-stores",
	Short: "Amazon Web Services HealthOmics variant stores and annotation stores will no longer be open to new customers starting November 7, 2025. If you would like to use variant stores or annotation stores, sign up prior to that date.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(omics_listAnnotationStoresCmd).Standalone()

	omics_listAnnotationStoresCmd.Flags().String("filter", "", "A filter to apply to the list.")
	omics_listAnnotationStoresCmd.Flags().String("ids", "", "IDs of stores to list.")
	omics_listAnnotationStoresCmd.Flags().String("max-results", "", "The maximum number of stores to return in one page of results.")
	omics_listAnnotationStoresCmd.Flags().String("next-token", "", "Specify the pagination token from a previous request to retrieve the next page of results.")
	omicsCmd.AddCommand(omics_listAnnotationStoresCmd)
}
