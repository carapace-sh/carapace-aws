package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_listTableOptimizerRunsCmd = &cobra.Command{
	Use:   "list-table-optimizer-runs",
	Short: "Lists the history of previous optimizer runs for a specific table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_listTableOptimizerRunsCmd).Standalone()

	glue_listTableOptimizerRunsCmd.Flags().String("catalog-id", "", "The Catalog ID of the table.")
	glue_listTableOptimizerRunsCmd.Flags().String("database-name", "", "The name of the database in the catalog in which the table resides.")
	glue_listTableOptimizerRunsCmd.Flags().String("max-results", "", "The maximum number of optimizer runs to return on each call.")
	glue_listTableOptimizerRunsCmd.Flags().String("next-token", "", "A continuation token, if this is a continuation call.")
	glue_listTableOptimizerRunsCmd.Flags().String("table-name", "", "The name of the table.")
	glue_listTableOptimizerRunsCmd.Flags().String("type", "", "The type of table optimizer.")
	glue_listTableOptimizerRunsCmd.MarkFlagRequired("catalog-id")
	glue_listTableOptimizerRunsCmd.MarkFlagRequired("database-name")
	glue_listTableOptimizerRunsCmd.MarkFlagRequired("table-name")
	glue_listTableOptimizerRunsCmd.MarkFlagRequired("type")
	glueCmd.AddCommand(glue_listTableOptimizerRunsCmd)
}
