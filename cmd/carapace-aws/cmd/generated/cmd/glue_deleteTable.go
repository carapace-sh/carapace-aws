package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_deleteTableCmd = &cobra.Command{
	Use:   "delete-table",
	Short: "Removes a table definition from the Data Catalog.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_deleteTableCmd).Standalone()

	glue_deleteTableCmd.Flags().String("catalog-id", "", "The ID of the Data Catalog where the table resides.")
	glue_deleteTableCmd.Flags().String("database-name", "", "The name of the catalog database in which the table resides.")
	glue_deleteTableCmd.Flags().String("name", "", "The name of the table to be deleted.")
	glue_deleteTableCmd.Flags().String("transaction-id", "", "The transaction ID at which to delete the table contents.")
	glue_deleteTableCmd.MarkFlagRequired("database-name")
	glue_deleteTableCmd.MarkFlagRequired("name")
	glueCmd.AddCommand(glue_deleteTableCmd)
}
