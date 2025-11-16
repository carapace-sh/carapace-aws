package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pricing_getPriceListFileUrlCmd = &cobra.Command{
	Use:   "get-price-list-file-url",
	Short: "***This feature is in preview release and is subject to change.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pricing_getPriceListFileUrlCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pricing_getPriceListFileUrlCmd).Standalone()

		pricing_getPriceListFileUrlCmd.Flags().String("file-format", "", "The format that you want to retrieve your Price List files in.")
		pricing_getPriceListFileUrlCmd.Flags().String("price-list-arn", "", "The unique identifier that maps to where your Price List files are located.")
		pricing_getPriceListFileUrlCmd.MarkFlagRequired("file-format")
		pricing_getPriceListFileUrlCmd.MarkFlagRequired("price-list-arn")
	})
	pricingCmd.AddCommand(pricing_getPriceListFileUrlCmd)
}
