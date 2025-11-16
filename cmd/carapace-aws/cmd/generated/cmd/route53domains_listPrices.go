package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53domains_listPricesCmd = &cobra.Command{
	Use:   "list-prices",
	Short: "Lists the following prices for either all the TLDs supported by Route\u00a053, or the specified TLD:\n\n- Registration\n- Transfer\n- Owner change\n- Domain renewal\n- Domain restoration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53domains_listPricesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53domains_listPricesCmd).Standalone()

		route53domains_listPricesCmd.Flags().String("marker", "", "For an initial request for a list of prices, omit this element.")
		route53domains_listPricesCmd.Flags().String("max-items", "", "Number of `Prices` to be returned.")
		route53domains_listPricesCmd.Flags().String("tld", "", "The TLD for which you want to receive the pricing information.")
	})
	route53domainsCmd.AddCommand(route53domains_listPricesCmd)
}
