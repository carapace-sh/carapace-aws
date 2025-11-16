package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pricing_listPriceListsCmd = &cobra.Command{
	Use:   "list-price-lists",
	Short: "***This feature is in preview release and is subject to change.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pricing_listPriceListsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pricing_listPriceListsCmd).Standalone()

		pricing_listPriceListsCmd.Flags().String("currency-code", "", "The three alphabetical character ISO-4217 currency code that the Price List files are denominated in.")
		pricing_listPriceListsCmd.Flags().String("effective-date", "", "The date that the Price List file prices are effective from.")
		pricing_listPriceListsCmd.Flags().String("max-results", "", "The maximum number of results to return in the response.")
		pricing_listPriceListsCmd.Flags().String("next-token", "", "The pagination token that indicates the next set of results that you want to retrieve.")
		pricing_listPriceListsCmd.Flags().String("region-code", "", "This is used to filter the Price List by Amazon Web Services Region.")
		pricing_listPriceListsCmd.Flags().String("service-code", "", "The service code or the Savings Plan service code for the attributes that you want to retrieve.")
		pricing_listPriceListsCmd.MarkFlagRequired("currency-code")
		pricing_listPriceListsCmd.MarkFlagRequired("effective-date")
		pricing_listPriceListsCmd.MarkFlagRequired("service-code")
	})
	pricingCmd.AddCommand(pricing_listPriceListsCmd)
}
