package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var taxsettings_batchPutTaxRegistrationCmd = &cobra.Command{
	Use:   "batch-put-tax-registration",
	Short: "Adds or updates tax registration for multiple accounts in batch.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(taxsettings_batchPutTaxRegistrationCmd).Standalone()

	taxsettings_batchPutTaxRegistrationCmd.Flags().String("account-ids", "", "List of unique account identifiers.")
	taxsettings_batchPutTaxRegistrationCmd.Flags().String("tax-registration-entry", "", "Your TRN information that will be stored to the accounts mentioned in `putEntries`.")
	taxsettings_batchPutTaxRegistrationCmd.MarkFlagRequired("account-ids")
	taxsettings_batchPutTaxRegistrationCmd.MarkFlagRequired("tax-registration-entry")
	taxsettingsCmd.AddCommand(taxsettings_batchPutTaxRegistrationCmd)
}
