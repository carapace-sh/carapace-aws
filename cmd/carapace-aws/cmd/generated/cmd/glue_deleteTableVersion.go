package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_deleteTableVersionCmd = &cobra.Command{
	Use:   "delete-table-version",
	Short: "Deletes a specified version of a table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_deleteTableVersionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_deleteTableVersionCmd).Standalone()

		glue_deleteTableVersionCmd.Flags().String("catalog-id", "", "The ID of the Data Catalog where the tables reside.")
		glue_deleteTableVersionCmd.Flags().String("database-name", "", "The database in the catalog in which the table resides.")
		glue_deleteTableVersionCmd.Flags().String("table-name", "", "The name of the table.")
		glue_deleteTableVersionCmd.Flags().String("version-id", "", "The ID of the table version to be deleted.")
		glue_deleteTableVersionCmd.MarkFlagRequired("database-name")
		glue_deleteTableVersionCmd.MarkFlagRequired("table-name")
		glue_deleteTableVersionCmd.MarkFlagRequired("version-id")
	})
	glueCmd.AddCommand(glue_deleteTableVersionCmd)
}
