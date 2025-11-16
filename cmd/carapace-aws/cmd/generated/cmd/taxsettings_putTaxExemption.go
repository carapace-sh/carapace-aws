package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var taxsettings_putTaxExemptionCmd = &cobra.Command{
	Use:   "put-tax-exemption",
	Short: "Adds the tax exemption for a single account or all accounts listed in a consolidated billing family.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(taxsettings_putTaxExemptionCmd).Standalone()

	taxsettings_putTaxExemptionCmd.Flags().String("account-ids", "", "The list of unique account identifiers.")
	taxsettings_putTaxExemptionCmd.Flags().String("authority", "", "")
	taxsettings_putTaxExemptionCmd.Flags().String("exemption-certificate", "", "")
	taxsettings_putTaxExemptionCmd.Flags().String("exemption-type", "", "The exemption type.")
	taxsettings_putTaxExemptionCmd.MarkFlagRequired("account-ids")
	taxsettings_putTaxExemptionCmd.MarkFlagRequired("authority")
	taxsettings_putTaxExemptionCmd.MarkFlagRequired("exemption-certificate")
	taxsettings_putTaxExemptionCmd.MarkFlagRequired("exemption-type")
	taxsettingsCmd.AddCommand(taxsettings_putTaxExemptionCmd)
}
