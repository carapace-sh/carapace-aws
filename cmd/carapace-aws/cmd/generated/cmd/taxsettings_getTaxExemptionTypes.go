package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var taxsettings_getTaxExemptionTypesCmd = &cobra.Command{
	Use:   "get-tax-exemption-types",
	Short: "Get supported tax exemption types.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(taxsettings_getTaxExemptionTypesCmd).Standalone()

	taxsettingsCmd.AddCommand(taxsettings_getTaxExemptionTypesCmd)
}
