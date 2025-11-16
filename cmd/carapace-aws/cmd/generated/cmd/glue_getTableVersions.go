package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_getTableVersionsCmd = &cobra.Command{
	Use:   "get-table-versions",
	Short: "Retrieves a list of strings that identify available versions of a specified table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_getTableVersionsCmd).Standalone()

	glue_getTableVersionsCmd.Flags().String("catalog-id", "", "The ID of the Data Catalog where the tables reside.")
	glue_getTableVersionsCmd.Flags().String("database-name", "", "The database in the catalog in which the table resides.")
	glue_getTableVersionsCmd.Flags().String("max-results", "", "The maximum number of table versions to return in one response.")
	glue_getTableVersionsCmd.Flags().String("next-token", "", "A continuation token, if this is not the first call.")
	glue_getTableVersionsCmd.Flags().String("table-name", "", "The name of the table.")
	glue_getTableVersionsCmd.MarkFlagRequired("database-name")
	glue_getTableVersionsCmd.MarkFlagRequired("table-name")
	glueCmd.AddCommand(glue_getTableVersionsCmd)
}
