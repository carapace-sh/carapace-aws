package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_getCatalogCmd = &cobra.Command{
	Use:   "get-catalog",
	Short: "The name of the Catalog to retrieve.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_getCatalogCmd).Standalone()

	glue_getCatalogCmd.Flags().String("catalog-id", "", "The ID of the parent catalog in which the catalog resides.")
	glue_getCatalogCmd.MarkFlagRequired("catalog-id")
	glueCmd.AddCommand(glue_getCatalogCmd)
}
