package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var athena_deleteDataCatalogCmd = &cobra.Command{
	Use:   "delete-data-catalog",
	Short: "Deletes a data catalog.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(athena_deleteDataCatalogCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(athena_deleteDataCatalogCmd).Standalone()

		athena_deleteDataCatalogCmd.Flags().Bool("delete-catalog-only", false, "Deletes the Athena Data Catalog.")
		athena_deleteDataCatalogCmd.Flags().String("name", "", "The name of the data catalog to delete.")
		athena_deleteDataCatalogCmd.Flags().Bool("no-delete-catalog-only", false, "Deletes the Athena Data Catalog.")
		athena_deleteDataCatalogCmd.MarkFlagRequired("name")
		athena_deleteDataCatalogCmd.Flag("no-delete-catalog-only").Hidden = true
	})
	athenaCmd.AddCommand(athena_deleteDataCatalogCmd)
}
