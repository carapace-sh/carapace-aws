package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_createCatalogCmd = &cobra.Command{
	Use:   "create-catalog",
	Short: "Creates a new catalog in the Glue Data Catalog.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_createCatalogCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_createCatalogCmd).Standalone()

		glue_createCatalogCmd.Flags().String("catalog-input", "", "A `CatalogInput` object that defines the metadata for the catalog.")
		glue_createCatalogCmd.Flags().String("name", "", "The name of the catalog to create.")
		glue_createCatalogCmd.Flags().String("tags", "", "A map array of key-value pairs, not more than 50 pairs.")
		glue_createCatalogCmd.MarkFlagRequired("catalog-input")
		glue_createCatalogCmd.MarkFlagRequired("name")
	})
	glueCmd.AddCommand(glue_createCatalogCmd)
}
