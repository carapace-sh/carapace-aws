package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_importCatalogToGlueCmd = &cobra.Command{
	Use:   "import-catalog-to-glue",
	Short: "Imports an existing Amazon Athena Data Catalog to Glue.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_importCatalogToGlueCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_importCatalogToGlueCmd).Standalone()

		glue_importCatalogToGlueCmd.Flags().String("catalog-id", "", "The ID of the catalog to import.")
	})
	glueCmd.AddCommand(glue_importCatalogToGlueCmd)
}
