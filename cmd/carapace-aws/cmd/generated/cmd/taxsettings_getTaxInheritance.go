package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var taxsettings_getTaxInheritanceCmd = &cobra.Command{
	Use:   "get-tax-inheritance",
	Short: "The get account tax inheritance status.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(taxsettings_getTaxInheritanceCmd).Standalone()

	taxsettingsCmd.AddCommand(taxsettings_getTaxInheritanceCmd)
}
