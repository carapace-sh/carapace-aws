package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var taxsettings_batchGetTaxExemptionsCmd = &cobra.Command{
	Use:   "batch-get-tax-exemptions",
	Short: "Get the active tax exemptions for a given list of accounts.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(taxsettings_batchGetTaxExemptionsCmd).Standalone()

	taxsettings_batchGetTaxExemptionsCmd.Flags().String("account-ids", "", "List of unique account identifiers.")
	taxsettings_batchGetTaxExemptionsCmd.MarkFlagRequired("account-ids")
	taxsettingsCmd.AddCommand(taxsettings_batchGetTaxExemptionsCmd)
}
