package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_deleteColumnStatisticsForPartitionCmd = &cobra.Command{
	Use:   "delete-column-statistics-for-partition",
	Short: "Delete the partition column statistics of a column.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_deleteColumnStatisticsForPartitionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_deleteColumnStatisticsForPartitionCmd).Standalone()

		glue_deleteColumnStatisticsForPartitionCmd.Flags().String("catalog-id", "", "The ID of the Data Catalog where the partitions in question reside.")
		glue_deleteColumnStatisticsForPartitionCmd.Flags().String("column-name", "", "Name of the column.")
		glue_deleteColumnStatisticsForPartitionCmd.Flags().String("database-name", "", "The name of the catalog database where the partitions reside.")
		glue_deleteColumnStatisticsForPartitionCmd.Flags().String("partition-values", "", "A list of partition values identifying the partition.")
		glue_deleteColumnStatisticsForPartitionCmd.Flags().String("table-name", "", "The name of the partitions' table.")
		glue_deleteColumnStatisticsForPartitionCmd.MarkFlagRequired("column-name")
		glue_deleteColumnStatisticsForPartitionCmd.MarkFlagRequired("database-name")
		glue_deleteColumnStatisticsForPartitionCmd.MarkFlagRequired("partition-values")
		glue_deleteColumnStatisticsForPartitionCmd.MarkFlagRequired("table-name")
	})
	glueCmd.AddCommand(glue_deleteColumnStatisticsForPartitionCmd)
}
