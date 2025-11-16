package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_createTableOptimizerCmd = &cobra.Command{
	Use:   "create-table-optimizer",
	Short: "Creates a new table optimizer for a specific function.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_createTableOptimizerCmd).Standalone()

	glue_createTableOptimizerCmd.Flags().String("catalog-id", "", "The Catalog ID of the table.")
	glue_createTableOptimizerCmd.Flags().String("database-name", "", "The name of the database in the catalog in which the table resides.")
	glue_createTableOptimizerCmd.Flags().String("table-name", "", "The name of the table.")
	glue_createTableOptimizerCmd.Flags().String("table-optimizer-configuration", "", "A `TableOptimizerConfiguration` object representing the configuration of a table optimizer.")
	glue_createTableOptimizerCmd.Flags().String("type", "", "The type of table optimizer.")
	glue_createTableOptimizerCmd.MarkFlagRequired("catalog-id")
	glue_createTableOptimizerCmd.MarkFlagRequired("database-name")
	glue_createTableOptimizerCmd.MarkFlagRequired("table-name")
	glue_createTableOptimizerCmd.MarkFlagRequired("table-optimizer-configuration")
	glue_createTableOptimizerCmd.MarkFlagRequired("type")
	glueCmd.AddCommand(glue_createTableOptimizerCmd)
}
