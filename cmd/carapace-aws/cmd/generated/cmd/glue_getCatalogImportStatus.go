package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_getCatalogImportStatusCmd = &cobra.Command{
	Use:   "get-catalog-import-status",
	Short: "Retrieves the status of a migration operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_getCatalogImportStatusCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_getCatalogImportStatusCmd).Standalone()

		glue_getCatalogImportStatusCmd.Flags().String("catalog-id", "", "The ID of the catalog to migrate.")
	})
	glueCmd.AddCommand(glue_getCatalogImportStatusCmd)
}
