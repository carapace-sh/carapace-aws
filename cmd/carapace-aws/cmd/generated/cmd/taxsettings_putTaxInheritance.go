package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var taxsettings_putTaxInheritanceCmd = &cobra.Command{
	Use:   "put-tax-inheritance",
	Short: "The updated tax inheritance status.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(taxsettings_putTaxInheritanceCmd).Standalone()

	taxsettings_putTaxInheritanceCmd.Flags().String("heritage-status", "", "The tax inheritance status.")
	taxsettingsCmd.AddCommand(taxsettings_putTaxInheritanceCmd)
}
