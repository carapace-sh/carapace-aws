package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var taxsettings_listSupplementalTaxRegistrationsCmd = &cobra.Command{
	Use:   "list-supplemental-tax-registrations",
	Short: "Retrieves supplemental tax registrations for a single account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(taxsettings_listSupplementalTaxRegistrationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(taxsettings_listSupplementalTaxRegistrationsCmd).Standalone()

		taxsettings_listSupplementalTaxRegistrationsCmd.Flags().String("max-results", "", "The number of `taxRegistrations` results you want in one response.")
		taxsettings_listSupplementalTaxRegistrationsCmd.Flags().String("next-token", "", "The token to retrieve the next set of results.")
	})
	taxsettingsCmd.AddCommand(taxsettings_listSupplementalTaxRegistrationsCmd)
}
