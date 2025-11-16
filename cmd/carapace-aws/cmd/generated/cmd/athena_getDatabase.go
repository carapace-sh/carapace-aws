package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var athena_getDatabaseCmd = &cobra.Command{
	Use:   "get-database",
	Short: "Returns a database object for the specified database and data catalog.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(athena_getDatabaseCmd).Standalone()

	athena_getDatabaseCmd.Flags().String("catalog-name", "", "The name of the data catalog that contains the database to return.")
	athena_getDatabaseCmd.Flags().String("database-name", "", "The name of the database to return.")
	athena_getDatabaseCmd.Flags().String("work-group", "", "The name of the workgroup for which the metadata is being fetched.")
	athena_getDatabaseCmd.MarkFlagRequired("catalog-name")
	athena_getDatabaseCmd.MarkFlagRequired("database-name")
	athenaCmd.AddCommand(athena_getDatabaseCmd)
}
