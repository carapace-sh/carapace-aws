package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var taxsettings_getTaxRegistrationCmd = &cobra.Command{
	Use:   "get-tax-registration",
	Short: "Retrieves tax registration for a single account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(taxsettings_getTaxRegistrationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(taxsettings_getTaxRegistrationCmd).Standalone()

		taxsettings_getTaxRegistrationCmd.Flags().String("account-id", "", "Your unique account identifier.")
	})
	taxsettingsCmd.AddCommand(taxsettings_getTaxRegistrationCmd)
}
