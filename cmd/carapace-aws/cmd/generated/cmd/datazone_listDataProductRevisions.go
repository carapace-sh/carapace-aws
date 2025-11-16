package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_listDataProductRevisionsCmd = &cobra.Command{
	Use:   "list-data-product-revisions",
	Short: "Lists data product revisions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_listDataProductRevisionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datazone_listDataProductRevisionsCmd).Standalone()

		datazone_listDataProductRevisionsCmd.Flags().String("domain-identifier", "", "The ID of the domain of the data product revisions that you want to list.")
		datazone_listDataProductRevisionsCmd.Flags().String("identifier", "", "The ID of the data product revision.")
		datazone_listDataProductRevisionsCmd.Flags().String("max-results", "", "The maximum number of asset filters to return in a single call to `ListDataProductRevisions`.")
		datazone_listDataProductRevisionsCmd.Flags().String("next-token", "", "When the number of data product revisions is greater than the default value for the `MaxResults` parameter, or if you explicitly specify a value for `MaxResults` that is less than the number of data product revisions, the response includes a pagination token named `NextToken`.")
		datazone_listDataProductRevisionsCmd.MarkFlagRequired("domain-identifier")
		datazone_listDataProductRevisionsCmd.MarkFlagRequired("identifier")
	})
	datazoneCmd.AddCommand(datazone_listDataProductRevisionsCmd)
}
