package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_createDatabaseCmd = &cobra.Command{
	Use:   "create-database",
	Short: "Creates a new database in a Data Catalog.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_createDatabaseCmd).Standalone()

	glue_createDatabaseCmd.Flags().String("catalog-id", "", "The ID of the Data Catalog in which to create the database.")
	glue_createDatabaseCmd.Flags().String("database-input", "", "The metadata for the database.")
	glue_createDatabaseCmd.Flags().String("tags", "", "The tags you assign to the database.")
	glue_createDatabaseCmd.MarkFlagRequired("database-input")
	glueCmd.AddCommand(glue_createDatabaseCmd)
}
