package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var taxsettings_batchDeleteTaxRegistrationCmd = &cobra.Command{
	Use:   "batch-delete-tax-registration",
	Short: "Deletes tax registration for multiple accounts in batch.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(taxsettings_batchDeleteTaxRegistrationCmd).Standalone()

	taxsettings_batchDeleteTaxRegistrationCmd.Flags().String("account-ids", "", "List of unique account identifiers.")
	taxsettings_batchDeleteTaxRegistrationCmd.MarkFlagRequired("account-ids")
	taxsettingsCmd.AddCommand(taxsettings_batchDeleteTaxRegistrationCmd)
}
