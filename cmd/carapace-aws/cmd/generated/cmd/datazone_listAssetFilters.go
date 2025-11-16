package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_listAssetFiltersCmd = &cobra.Command{
	Use:   "list-asset-filters",
	Short: "Lists asset filters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_listAssetFiltersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datazone_listAssetFiltersCmd).Standalone()

		datazone_listAssetFiltersCmd.Flags().String("asset-identifier", "", "The ID of the data asset.")
		datazone_listAssetFiltersCmd.Flags().String("domain-identifier", "", "The ID of the domain where you want to list asset filters.")
		datazone_listAssetFiltersCmd.Flags().String("max-results", "", "The maximum number of asset filters to return in a single call to `ListAssetFilters`.")
		datazone_listAssetFiltersCmd.Flags().String("next-token", "", "When the number of asset filters is greater than the default value for the `MaxResults` parameter, or if you explicitly specify a value for `MaxResults` that is less than the number of asset filters, the response includes a pagination token named `NextToken`.")
		datazone_listAssetFiltersCmd.Flags().String("status", "", "The status of the asset filter.")
		datazone_listAssetFiltersCmd.MarkFlagRequired("asset-identifier")
		datazone_listAssetFiltersCmd.MarkFlagRequired("domain-identifier")
	})
	datazoneCmd.AddCommand(datazone_listAssetFiltersCmd)
}
