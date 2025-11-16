package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var marketplaceCatalogCmd = &cobra.Command{
	Use:   "marketplace-catalog",
	Short: "Catalog API actions allow you to manage your entities through list, describe, and update capabilities.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(marketplaceCatalogCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(marketplaceCatalogCmd).Standalone()

	})
	rootCmd.AddCommand(marketplaceCatalogCmd)
}
