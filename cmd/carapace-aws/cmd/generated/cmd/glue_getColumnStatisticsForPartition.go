package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_getColumnStatisticsForPartitionCmd = &cobra.Command{
	Use:   "get-column-statistics-for-partition",
	Short: "Retrieves partition statistics of columns.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_getColumnStatisticsForPartitionCmd).Standalone()

	glue_getColumnStatisticsForPartitionCmd.Flags().String("catalog-id", "", "The ID of the Data Catalog where the partitions in question reside.")
	glue_getColumnStatisticsForPartitionCmd.Flags().String("column-names", "", "A list of the column names.")
	glue_getColumnStatisticsForPartitionCmd.Flags().String("database-name", "", "The name of the catalog database where the partitions reside.")
	glue_getColumnStatisticsForPartitionCmd.Flags().String("partition-values", "", "A list of partition values identifying the partition.")
	glue_getColumnStatisticsForPartitionCmd.Flags().String("table-name", "", "The name of the partitions' table.")
	glue_getColumnStatisticsForPartitionCmd.MarkFlagRequired("column-names")
	glue_getColumnStatisticsForPartitionCmd.MarkFlagRequired("database-name")
	glue_getColumnStatisticsForPartitionCmd.MarkFlagRequired("partition-values")
	glue_getColumnStatisticsForPartitionCmd.MarkFlagRequired("table-name")
	glueCmd.AddCommand(glue_getColumnStatisticsForPartitionCmd)
}
