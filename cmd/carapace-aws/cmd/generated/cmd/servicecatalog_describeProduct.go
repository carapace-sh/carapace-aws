package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_describeProductCmd = &cobra.Command{
	Use:   "describe-product",
	Short: "Gets information about the specified product.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_describeProductCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(servicecatalog_describeProductCmd).Standalone()

		servicecatalog_describeProductCmd.Flags().String("accept-language", "", "The language code.")
		servicecatalog_describeProductCmd.Flags().String("id", "", "The product identifier.")
		servicecatalog_describeProductCmd.Flags().String("name", "", "The product name.")
	})
	servicecatalogCmd.AddCommand(servicecatalog_describeProductCmd)
}
