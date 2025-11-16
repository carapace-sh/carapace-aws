package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_batchDeleteTableVersionCmd = &cobra.Command{
	Use:   "batch-delete-table-version",
	Short: "Deletes a specified batch of versions of a table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_batchDeleteTableVersionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_batchDeleteTableVersionCmd).Standalone()

		glue_batchDeleteTableVersionCmd.Flags().String("catalog-id", "", "The ID of the Data Catalog where the tables reside.")
		glue_batchDeleteTableVersionCmd.Flags().String("database-name", "", "The database in the catalog in which the table resides.")
		glue_batchDeleteTableVersionCmd.Flags().String("table-name", "", "The name of the table.")
		glue_batchDeleteTableVersionCmd.Flags().String("version-ids", "", "A list of the IDs of versions to be deleted.")
		glue_batchDeleteTableVersionCmd.MarkFlagRequired("database-name")
		glue_batchDeleteTableVersionCmd.MarkFlagRequired("table-name")
		glue_batchDeleteTableVersionCmd.MarkFlagRequired("version-ids")
	})
	glueCmd.AddCommand(glue_batchDeleteTableVersionCmd)
}
