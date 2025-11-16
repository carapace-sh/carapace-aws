package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pricing_getProductsCmd = &cobra.Command{
	Use:   "get-products",
	Short: "Returns a list of all products that match the filter criteria.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pricing_getProductsCmd).Standalone()

	pricing_getProductsCmd.Flags().String("filters", "", "The list of filters that limit the returned products.")
	pricing_getProductsCmd.Flags().String("format-version", "", "The format version that you want the response to be in.")
	pricing_getProductsCmd.Flags().String("max-results", "", "The maximum number of results to return in the response.")
	pricing_getProductsCmd.Flags().String("next-token", "", "The pagination token that indicates the next set of results that you want to retrieve.")
	pricing_getProductsCmd.Flags().String("service-code", "", "The code for the service whose products you want to retrieve.")
	pricing_getProductsCmd.MarkFlagRequired("service-code")
	pricingCmd.AddCommand(pricing_getProductsCmd)
}
