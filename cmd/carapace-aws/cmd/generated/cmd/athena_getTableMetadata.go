package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var athena_getTableMetadataCmd = &cobra.Command{
	Use:   "get-table-metadata",
	Short: "Returns table metadata for the specified catalog, database, and table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(athena_getTableMetadataCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(athena_getTableMetadataCmd).Standalone()

		athena_getTableMetadataCmd.Flags().String("catalog-name", "", "The name of the data catalog that contains the database and table metadata to return.")
		athena_getTableMetadataCmd.Flags().String("database-name", "", "The name of the database that contains the table metadata to return.")
		athena_getTableMetadataCmd.Flags().String("table-name", "", "The name of the table for which metadata is returned.")
		athena_getTableMetadataCmd.Flags().String("work-group", "", "The name of the workgroup for which the metadata is being fetched.")
		athena_getTableMetadataCmd.MarkFlagRequired("catalog-name")
		athena_getTableMetadataCmd.MarkFlagRequired("database-name")
		athena_getTableMetadataCmd.MarkFlagRequired("table-name")
	})
	athenaCmd.AddCommand(athena_getTableMetadataCmd)
}
