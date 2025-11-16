package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_deleteCatalogCmd = &cobra.Command{
	Use:   "delete-catalog",
	Short: "Removes the specified catalog from the Glue Data Catalog.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_deleteCatalogCmd).Standalone()

	glue_deleteCatalogCmd.Flags().String("catalog-id", "", "The ID of the catalog.")
	glue_deleteCatalogCmd.MarkFlagRequired("catalog-id")
	glueCmd.AddCommand(glue_deleteCatalogCmd)
}
