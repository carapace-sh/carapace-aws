package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var taxsettingsCmd = &cobra.Command{
	Use:   "taxsettings",
	Short: "You can use the tax setting API to programmatically set, modify, and delete the tax registration number (TRN), associated business legal name, and address (Collectively referred to as \"TRN information\").",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(taxsettingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(taxsettingsCmd).Standalone()

	})
	rootCmd.AddCommand(taxsettingsCmd)
}
