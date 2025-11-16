package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_updateCatalogCmd = &cobra.Command{
	Use:   "update-catalog",
	Short: "Updates an existing catalog's properties in the Glue Data Catalog.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_updateCatalogCmd).Standalone()

	glue_updateCatalogCmd.Flags().String("catalog-id", "", "The ID of the catalog.")
	glue_updateCatalogCmd.Flags().String("catalog-input", "", "A `CatalogInput` object specifying the new properties of an existing catalog.")
	glue_updateCatalogCmd.MarkFlagRequired("catalog-id")
	glue_updateCatalogCmd.MarkFlagRequired("catalog-input")
	glueCmd.AddCommand(glue_updateCatalogCmd)
}
