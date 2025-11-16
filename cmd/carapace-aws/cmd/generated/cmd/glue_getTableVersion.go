package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_getTableVersionCmd = &cobra.Command{
	Use:   "get-table-version",
	Short: "Retrieves a specified version of a table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_getTableVersionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_getTableVersionCmd).Standalone()

		glue_getTableVersionCmd.Flags().String("catalog-id", "", "The ID of the Data Catalog where the tables reside.")
		glue_getTableVersionCmd.Flags().String("database-name", "", "The database in the catalog in which the table resides.")
		glue_getTableVersionCmd.Flags().String("table-name", "", "The name of the table.")
		glue_getTableVersionCmd.Flags().String("version-id", "", "The ID value of the table version to be retrieved.")
		glue_getTableVersionCmd.MarkFlagRequired("database-name")
		glue_getTableVersionCmd.MarkFlagRequired("table-name")
	})
	glueCmd.AddCommand(glue_getTableVersionCmd)
}
