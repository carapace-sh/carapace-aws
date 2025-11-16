package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var taxsettings_deleteSupplementalTaxRegistrationCmd = &cobra.Command{
	Use:   "delete-supplemental-tax-registration",
	Short: "Deletes a supplemental tax registration for a single account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(taxsettings_deleteSupplementalTaxRegistrationCmd).Standalone()

	taxsettings_deleteSupplementalTaxRegistrationCmd.Flags().String("authority-id", "", "The unique authority Id for the supplemental TRN information that needs to be deleted.")
	taxsettings_deleteSupplementalTaxRegistrationCmd.MarkFlagRequired("authority-id")
	taxsettingsCmd.AddCommand(taxsettings_deleteSupplementalTaxRegistrationCmd)
}
