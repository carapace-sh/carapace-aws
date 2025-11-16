package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_deleteTableOptimizerCmd = &cobra.Command{
	Use:   "delete-table-optimizer",
	Short: "Deletes an optimizer and all associated metadata for a table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_deleteTableOptimizerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_deleteTableOptimizerCmd).Standalone()

		glue_deleteTableOptimizerCmd.Flags().String("catalog-id", "", "The Catalog ID of the table.")
		glue_deleteTableOptimizerCmd.Flags().String("database-name", "", "The name of the database in the catalog in which the table resides.")
		glue_deleteTableOptimizerCmd.Flags().String("table-name", "", "The name of the table.")
		glue_deleteTableOptimizerCmd.Flags().String("type", "", "The type of table optimizer.")
		glue_deleteTableOptimizerCmd.MarkFlagRequired("catalog-id")
		glue_deleteTableOptimizerCmd.MarkFlagRequired("database-name")
		glue_deleteTableOptimizerCmd.MarkFlagRequired("table-name")
		glue_deleteTableOptimizerCmd.MarkFlagRequired("type")
	})
	glueCmd.AddCommand(glue_deleteTableOptimizerCmd)
}
