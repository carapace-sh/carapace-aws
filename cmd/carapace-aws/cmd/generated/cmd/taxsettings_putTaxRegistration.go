package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var taxsettings_putTaxRegistrationCmd = &cobra.Command{
	Use:   "put-tax-registration",
	Short: "Adds or updates tax registration for a single account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(taxsettings_putTaxRegistrationCmd).Standalone()

	taxsettings_putTaxRegistrationCmd.Flags().String("account-id", "", "Your unique account identifier.")
	taxsettings_putTaxRegistrationCmd.Flags().String("tax-registration-entry", "", "Your TRN information that will be stored to the account mentioned in `accountId`.")
	taxsettings_putTaxRegistrationCmd.MarkFlagRequired("tax-registration-entry")
	taxsettingsCmd.AddCommand(taxsettings_putTaxRegistrationCmd)
}
