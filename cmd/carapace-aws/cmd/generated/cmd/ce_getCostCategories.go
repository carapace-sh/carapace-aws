package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ce_getCostCategoriesCmd = &cobra.Command{
	Use:   "get-cost-categories",
	Short: "Retrieves an array of Cost Category names and values incurred cost.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ce_getCostCategoriesCmd).Standalone()

	ce_getCostCategoriesCmd.Flags().String("billing-view-arn", "", "The Amazon Resource Name (ARN) that uniquely identifies a specific billing view.")
	ce_getCostCategoriesCmd.Flags().String("cost-category-name", "", "")
	ce_getCostCategoriesCmd.Flags().String("filter", "", "")
	ce_getCostCategoriesCmd.Flags().String("max-results", "", "This field is only used when the `SortBy` value is provided in the request.")
	ce_getCostCategoriesCmd.Flags().String("next-page-token", "", "If the number of objects that are still available for retrieval exceeds the quota, Amazon Web Services returns a NextPageToken value in the response.")
	ce_getCostCategoriesCmd.Flags().String("search-string", "", "The value that you want to search the filter values for.")
	ce_getCostCategoriesCmd.Flags().String("sort-by", "", "The value that you sort the data by.")
	ce_getCostCategoriesCmd.Flags().String("time-period", "", "")
	ce_getCostCategoriesCmd.MarkFlagRequired("time-period")
	ceCmd.AddCommand(ce_getCostCategoriesCmd)
}
