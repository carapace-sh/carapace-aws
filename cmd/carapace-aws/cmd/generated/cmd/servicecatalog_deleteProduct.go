package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_deleteProductCmd = &cobra.Command{
	Use:   "delete-product",
	Short: "Deletes the specified product.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_deleteProductCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(servicecatalog_deleteProductCmd).Standalone()

		servicecatalog_deleteProductCmd.Flags().String("accept-language", "", "The language code.")
		servicecatalog_deleteProductCmd.Flags().String("id", "", "The product identifier.")
		servicecatalog_deleteProductCmd.MarkFlagRequired("id")
	})
	servicecatalogCmd.AddCommand(servicecatalog_deleteProductCmd)
}
