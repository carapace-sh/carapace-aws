package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var outposts_listSitesCmd = &cobra.Command{
	Use:   "list-sites",
	Short: "Lists the Outpost sites for your Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(outposts_listSitesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(outposts_listSitesCmd).Standalone()

		outposts_listSitesCmd.Flags().String("max-results", "", "")
		outposts_listSitesCmd.Flags().String("next-token", "", "")
		outposts_listSitesCmd.Flags().String("operating-address-city-filter", "", "Filters the results by city.")
		outposts_listSitesCmd.Flags().String("operating-address-country-code-filter", "", "Filters the results by country code.")
		outposts_listSitesCmd.Flags().String("operating-address-state-or-region-filter", "", "Filters the results by state or region.")
	})
	outpostsCmd.AddCommand(outposts_listSitesCmd)
}
