package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var taxsettings_putSupplementalTaxRegistrationCmd = &cobra.Command{
	Use:   "put-supplemental-tax-registration",
	Short: "Stores supplemental tax registration for a single account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(taxsettings_putSupplementalTaxRegistrationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(taxsettings_putSupplementalTaxRegistrationCmd).Standalone()

		taxsettings_putSupplementalTaxRegistrationCmd.Flags().String("tax-registration-entry", "", "The supplemental TRN information that will be stored for the caller account ID.")
		taxsettings_putSupplementalTaxRegistrationCmd.MarkFlagRequired("tax-registration-entry")
	})
	taxsettingsCmd.AddCommand(taxsettings_putSupplementalTaxRegistrationCmd)
}
