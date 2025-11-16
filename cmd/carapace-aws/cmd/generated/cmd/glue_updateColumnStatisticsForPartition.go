package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_updateColumnStatisticsForPartitionCmd = &cobra.Command{
	Use:   "update-column-statistics-for-partition",
	Short: "Creates or updates partition statistics of columns.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_updateColumnStatisticsForPartitionCmd).Standalone()

	glue_updateColumnStatisticsForPartitionCmd.Flags().String("catalog-id", "", "The ID of the Data Catalog where the partitions in question reside.")
	glue_updateColumnStatisticsForPartitionCmd.Flags().String("column-statistics-list", "", "A list of the column statistics.")
	glue_updateColumnStatisticsForPartitionCmd.Flags().String("database-name", "", "The name of the catalog database where the partitions reside.")
	glue_updateColumnStatisticsForPartitionCmd.Flags().String("partition-values", "", "A list of partition values identifying the partition.")
	glue_updateColumnStatisticsForPartitionCmd.Flags().String("table-name", "", "The name of the partitions' table.")
	glue_updateColumnStatisticsForPartitionCmd.MarkFlagRequired("column-statistics-list")
	glue_updateColumnStatisticsForPartitionCmd.MarkFlagRequired("database-name")
	glue_updateColumnStatisticsForPartitionCmd.MarkFlagRequired("partition-values")
	glue_updateColumnStatisticsForPartitionCmd.MarkFlagRequired("table-name")
	glueCmd.AddCommand(glue_updateColumnStatisticsForPartitionCmd)
}
