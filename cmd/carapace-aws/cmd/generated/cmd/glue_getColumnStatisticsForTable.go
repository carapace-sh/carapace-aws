package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_getColumnStatisticsForTableCmd = &cobra.Command{
	Use:   "get-column-statistics-for-table",
	Short: "Retrieves table statistics of columns.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_getColumnStatisticsForTableCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_getColumnStatisticsForTableCmd).Standalone()

		glue_getColumnStatisticsForTableCmd.Flags().String("catalog-id", "", "The ID of the Data Catalog where the partitions in question reside.")
		glue_getColumnStatisticsForTableCmd.Flags().String("column-names", "", "A list of the column names.")
		glue_getColumnStatisticsForTableCmd.Flags().String("database-name", "", "The name of the catalog database where the partitions reside.")
		glue_getColumnStatisticsForTableCmd.Flags().String("table-name", "", "The name of the partitions' table.")
		glue_getColumnStatisticsForTableCmd.MarkFlagRequired("column-names")
		glue_getColumnStatisticsForTableCmd.MarkFlagRequired("database-name")
		glue_getColumnStatisticsForTableCmd.MarkFlagRequired("table-name")
	})
	glueCmd.AddCommand(glue_getColumnStatisticsForTableCmd)
}
