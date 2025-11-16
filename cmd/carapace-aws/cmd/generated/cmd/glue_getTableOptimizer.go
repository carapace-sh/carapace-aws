package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_getTableOptimizerCmd = &cobra.Command{
	Use:   "get-table-optimizer",
	Short: "Returns the configuration of all optimizers associated with a specified table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_getTableOptimizerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_getTableOptimizerCmd).Standalone()

		glue_getTableOptimizerCmd.Flags().String("catalog-id", "", "The Catalog ID of the table.")
		glue_getTableOptimizerCmd.Flags().String("database-name", "", "The name of the database in the catalog in which the table resides.")
		glue_getTableOptimizerCmd.Flags().String("table-name", "", "The name of the table.")
		glue_getTableOptimizerCmd.Flags().String("type", "", "The type of table optimizer.")
		glue_getTableOptimizerCmd.MarkFlagRequired("catalog-id")
		glue_getTableOptimizerCmd.MarkFlagRequired("database-name")
		glue_getTableOptimizerCmd.MarkFlagRequired("table-name")
		glue_getTableOptimizerCmd.MarkFlagRequired("type")
	})
	glueCmd.AddCommand(glue_getTableOptimizerCmd)
}
