package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_getDatabaseCmd = &cobra.Command{
	Use:   "get-database",
	Short: "Retrieves the definition of a specified database.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_getDatabaseCmd).Standalone()

	glue_getDatabaseCmd.Flags().String("catalog-id", "", "The ID of the Data Catalog in which the database resides.")
	glue_getDatabaseCmd.Flags().String("name", "", "The name of the database to retrieve.")
	glue_getDatabaseCmd.MarkFlagRequired("name")
	glueCmd.AddCommand(glue_getDatabaseCmd)
}
