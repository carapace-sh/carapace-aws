package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_describeProductAsAdminCmd = &cobra.Command{
	Use:   "describe-product-as-admin",
	Short: "Gets information about the specified product.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_describeProductAsAdminCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(servicecatalog_describeProductAsAdminCmd).Standalone()

		servicecatalog_describeProductAsAdminCmd.Flags().String("accept-language", "", "The language code.")
		servicecatalog_describeProductAsAdminCmd.Flags().String("id", "", "The product identifier.")
		servicecatalog_describeProductAsAdminCmd.Flags().String("name", "", "The product name.")
		servicecatalog_describeProductAsAdminCmd.Flags().String("source-portfolio-id", "", "The unique identifier of the shared portfolio that the specified product is associated with.")
	})
	servicecatalogCmd.AddCommand(servicecatalog_describeProductAsAdminCmd)
}
