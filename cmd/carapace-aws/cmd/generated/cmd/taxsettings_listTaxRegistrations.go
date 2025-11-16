package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var taxsettings_listTaxRegistrationsCmd = &cobra.Command{
	Use:   "list-tax-registrations",
	Short: "Retrieves the tax registration of accounts listed in a consolidated billing family.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(taxsettings_listTaxRegistrationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(taxsettings_listTaxRegistrationsCmd).Standalone()

		taxsettings_listTaxRegistrationsCmd.Flags().String("max-results", "", "Number of `accountDetails` results you want in one response.")
		taxsettings_listTaxRegistrationsCmd.Flags().String("next-token", "", "The token to retrieve the next set of results.")
	})
	taxsettingsCmd.AddCommand(taxsettings_listTaxRegistrationsCmd)
}
