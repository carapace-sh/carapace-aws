package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_updateDatabaseCmd = &cobra.Command{
	Use:   "update-database",
	Short: "Updates an existing database definition in a Data Catalog.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_updateDatabaseCmd).Standalone()

	glue_updateDatabaseCmd.Flags().String("catalog-id", "", "The ID of the Data Catalog in which the metadata database resides.")
	glue_updateDatabaseCmd.Flags().String("database-input", "", "A `DatabaseInput` object specifying the new definition of the metadata database in the catalog.")
	glue_updateDatabaseCmd.Flags().String("name", "", "The name of the database to update in the catalog.")
	glue_updateDatabaseCmd.MarkFlagRequired("database-input")
	glue_updateDatabaseCmd.MarkFlagRequired("name")
	glueCmd.AddCommand(glue_updateDatabaseCmd)
}
