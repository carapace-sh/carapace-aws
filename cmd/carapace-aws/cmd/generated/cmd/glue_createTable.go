package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_createTableCmd = &cobra.Command{
	Use:   "create-table",
	Short: "Creates a new table definition in the Data Catalog.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_createTableCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_createTableCmd).Standalone()

		glue_createTableCmd.Flags().String("catalog-id", "", "The ID of the Data Catalog in which to create the `Table`.")
		glue_createTableCmd.Flags().String("database-name", "", "The catalog database in which to create the new table.")
		glue_createTableCmd.Flags().String("name", "", "The unique identifier for the table within the specified database that will be created in the Glue Data Catalog.")
		glue_createTableCmd.Flags().String("open-table-format-input", "", "Specifies an `OpenTableFormatInput` structure when creating an open format table.")
		glue_createTableCmd.Flags().String("partition-indexes", "", "A list of partition indexes, `PartitionIndex` structures, to create in the table.")
		glue_createTableCmd.Flags().String("table-input", "", "The `TableInput` object that defines the metadata table to create in the catalog.")
		glue_createTableCmd.Flags().String("transaction-id", "", "The ID of the transaction.")
		glue_createTableCmd.MarkFlagRequired("database-name")
	})
	glueCmd.AddCommand(glue_createTableCmd)
}
