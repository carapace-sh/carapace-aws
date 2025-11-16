package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_deleteDatabaseCmd = &cobra.Command{
	Use:   "delete-database",
	Short: "Removes a specified database from a Data Catalog.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_deleteDatabaseCmd).Standalone()

	glue_deleteDatabaseCmd.Flags().String("catalog-id", "", "The ID of the Data Catalog in which the database resides.")
	glue_deleteDatabaseCmd.Flags().String("name", "", "The name of the database to delete.")
	glue_deleteDatabaseCmd.MarkFlagRequired("name")
	glueCmd.AddCommand(glue_deleteDatabaseCmd)
}
