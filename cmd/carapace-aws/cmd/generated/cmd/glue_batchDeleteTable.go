package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_batchDeleteTableCmd = &cobra.Command{
	Use:   "batch-delete-table",
	Short: "Deletes multiple tables at once.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_batchDeleteTableCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_batchDeleteTableCmd).Standalone()

		glue_batchDeleteTableCmd.Flags().String("catalog-id", "", "The ID of the Data Catalog where the table resides.")
		glue_batchDeleteTableCmd.Flags().String("database-name", "", "The name of the catalog database in which the tables to delete reside.")
		glue_batchDeleteTableCmd.Flags().String("tables-to-delete", "", "A list of the table to delete.")
		glue_batchDeleteTableCmd.Flags().String("transaction-id", "", "The transaction ID at which to delete the table contents.")
		glue_batchDeleteTableCmd.MarkFlagRequired("database-name")
		glue_batchDeleteTableCmd.MarkFlagRequired("tables-to-delete")
	})
	glueCmd.AddCommand(glue_batchDeleteTableCmd)
}
