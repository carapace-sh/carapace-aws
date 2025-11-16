package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var omics_listAnnotationStoreVersionsCmd = &cobra.Command{
	Use:   "list-annotation-store-versions",
	Short: "Lists the versions of an annotation store.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(omics_listAnnotationStoreVersionsCmd).Standalone()

	omics_listAnnotationStoreVersionsCmd.Flags().String("filter", "", "A filter to apply to the list of annotation store versions.")
	omics_listAnnotationStoreVersionsCmd.Flags().String("max-results", "", "The maximum number of annotation store versions to return in one page of results.")
	omics_listAnnotationStoreVersionsCmd.Flags().String("name", "", "The name of an annotation store.")
	omics_listAnnotationStoreVersionsCmd.Flags().String("next-token", "", "Specifies the pagination token from a previous request to retrieve the next page of results.")
	omics_listAnnotationStoreVersionsCmd.MarkFlagRequired("name")
	omicsCmd.AddCommand(omics_listAnnotationStoreVersionsCmd)
}
