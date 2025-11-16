package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_describeProductViewCmd = &cobra.Command{
	Use:   "describe-product-view",
	Short: "Gets information about the specified product.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_describeProductViewCmd).Standalone()

	servicecatalog_describeProductViewCmd.Flags().String("accept-language", "", "The language code.")
	servicecatalog_describeProductViewCmd.Flags().String("id", "", "The product view identifier.")
	servicecatalog_describeProductViewCmd.MarkFlagRequired("id")
	servicecatalogCmd.AddCommand(servicecatalog_describeProductViewCmd)
}
