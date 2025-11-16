package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_updateColumnStatisticsForTableCmd = &cobra.Command{
	Use:   "update-column-statistics-for-table",
	Short: "Creates or updates table statistics of columns.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_updateColumnStatisticsForTableCmd).Standalone()

	glue_updateColumnStatisticsForTableCmd.Flags().String("catalog-id", "", "The ID of the Data Catalog where the partitions in question reside.")
	glue_updateColumnStatisticsForTableCmd.Flags().String("column-statistics-list", "", "A list of the column statistics.")
	glue_updateColumnStatisticsForTableCmd.Flags().String("database-name", "", "The name of the catalog database where the partitions reside.")
	glue_updateColumnStatisticsForTableCmd.Flags().String("table-name", "", "The name of the partitions' table.")
	glue_updateColumnStatisticsForTableCmd.MarkFlagRequired("column-statistics-list")
	glue_updateColumnStatisticsForTableCmd.MarkFlagRequired("database-name")
	glue_updateColumnStatisticsForTableCmd.MarkFlagRequired("table-name")
	glueCmd.AddCommand(glue_updateColumnStatisticsForTableCmd)
}
