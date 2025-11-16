package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_updateTableOptimizerCmd = &cobra.Command{
	Use:   "update-table-optimizer",
	Short: "Updates the configuration for an existing table optimizer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_updateTableOptimizerCmd).Standalone()

	glue_updateTableOptimizerCmd.Flags().String("catalog-id", "", "The Catalog ID of the table.")
	glue_updateTableOptimizerCmd.Flags().String("database-name", "", "The name of the database in the catalog in which the table resides.")
	glue_updateTableOptimizerCmd.Flags().String("table-name", "", "The name of the table.")
	glue_updateTableOptimizerCmd.Flags().String("table-optimizer-configuration", "", "A `TableOptimizerConfiguration` object representing the configuration of a table optimizer.")
	glue_updateTableOptimizerCmd.Flags().String("type", "", "The type of table optimizer.")
	glue_updateTableOptimizerCmd.MarkFlagRequired("catalog-id")
	glue_updateTableOptimizerCmd.MarkFlagRequired("database-name")
	glue_updateTableOptimizerCmd.MarkFlagRequired("table-name")
	glue_updateTableOptimizerCmd.MarkFlagRequired("table-optimizer-configuration")
	glue_updateTableOptimizerCmd.MarkFlagRequired("type")
	glueCmd.AddCommand(glue_updateTableOptimizerCmd)
}
