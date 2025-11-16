package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var taxsettings_listTaxExemptionsCmd = &cobra.Command{
	Use:   "list-tax-exemptions",
	Short: "Retrieves the tax exemption of accounts listed in a consolidated billing family.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(taxsettings_listTaxExemptionsCmd).Standalone()

	taxsettings_listTaxExemptionsCmd.Flags().String("max-results", "", "The number of results you want in one response.")
	taxsettings_listTaxExemptionsCmd.Flags().String("next-token", "", "The token to retrieve the next set of results.")
	taxsettingsCmd.AddCommand(taxsettings_listTaxExemptionsCmd)
}
