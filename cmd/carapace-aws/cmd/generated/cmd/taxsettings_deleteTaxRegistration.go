package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var taxsettings_deleteTaxRegistrationCmd = &cobra.Command{
	Use:   "delete-tax-registration",
	Short: "Deletes tax registration for a single account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(taxsettings_deleteTaxRegistrationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(taxsettings_deleteTaxRegistrationCmd).Standalone()

		taxsettings_deleteTaxRegistrationCmd.Flags().String("account-id", "", "Unique account identifier for the TRN information that needs to be deleted.")
	})
	taxsettingsCmd.AddCommand(taxsettings_deleteTaxRegistrationCmd)
}
