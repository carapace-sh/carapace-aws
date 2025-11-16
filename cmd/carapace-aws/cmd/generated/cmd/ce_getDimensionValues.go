package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ce_getDimensionValuesCmd = &cobra.Command{
	Use:   "get-dimension-values",
	Short: "Retrieves all available filter values for a specified filter over a period of time.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ce_getDimensionValuesCmd).Standalone()

	ce_getDimensionValuesCmd.Flags().String("billing-view-arn", "", "The Amazon Resource Name (ARN) that uniquely identifies a specific billing view.")
	ce_getDimensionValuesCmd.Flags().String("context", "", "The context for the call to `GetDimensionValues`.")
	ce_getDimensionValuesCmd.Flags().String("dimension", "", "The name of the dimension.")
	ce_getDimensionValuesCmd.Flags().String("filter", "", "")
	ce_getDimensionValuesCmd.Flags().String("max-results", "", "This field is only used when SortBy is provided in the request.")
	ce_getDimensionValuesCmd.Flags().String("next-page-token", "", "The token to retrieve the next set of results.")
	ce_getDimensionValuesCmd.Flags().String("search-string", "", "The value that you want to search the filter values for.")
	ce_getDimensionValuesCmd.Flags().String("sort-by", "", "The value that you want to sort the data by.")
	ce_getDimensionValuesCmd.Flags().String("time-period", "", "The start date and end date for retrieving the dimension values.")
	ce_getDimensionValuesCmd.MarkFlagRequired("dimension")
	ce_getDimensionValuesCmd.MarkFlagRequired("time-period")
	ceCmd.AddCommand(ce_getDimensionValuesCmd)
}
