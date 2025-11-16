package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_deleteColumnStatisticsForTableCmd = &cobra.Command{
	Use:   "delete-column-statistics-for-table",
	Short: "Retrieves table statistics of columns.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_deleteColumnStatisticsForTableCmd).Standalone()

	glue_deleteColumnStatisticsForTableCmd.Flags().String("catalog-id", "", "The ID of the Data Catalog where the partitions in question reside.")
	glue_deleteColumnStatisticsForTableCmd.Flags().String("column-name", "", "The name of the column.")
	glue_deleteColumnStatisticsForTableCmd.Flags().String("database-name", "", "The name of the catalog database where the partitions reside.")
	glue_deleteColumnStatisticsForTableCmd.Flags().String("table-name", "", "The name of the partitions' table.")
	glue_deleteColumnStatisticsForTableCmd.MarkFlagRequired("column-name")
	glue_deleteColumnStatisticsForTableCmd.MarkFlagRequired("database-name")
	glue_deleteColumnStatisticsForTableCmd.MarkFlagRequired("table-name")
	glueCmd.AddCommand(glue_deleteColumnStatisticsForTableCmd)
}
