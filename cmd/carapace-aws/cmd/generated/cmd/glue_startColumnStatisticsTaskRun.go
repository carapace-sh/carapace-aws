package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_startColumnStatisticsTaskRunCmd = &cobra.Command{
	Use:   "start-column-statistics-task-run",
	Short: "Starts a column statistics task run, for a specified table and columns.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_startColumnStatisticsTaskRunCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_startColumnStatisticsTaskRunCmd).Standalone()

		glue_startColumnStatisticsTaskRunCmd.Flags().String("catalog-id", "", "The ID of the Data Catalog where the table reside.")
		glue_startColumnStatisticsTaskRunCmd.Flags().String("column-name-list", "", "A list of the column names to generate statistics.")
		glue_startColumnStatisticsTaskRunCmd.Flags().String("database-name", "", "The name of the database where the table resides.")
		glue_startColumnStatisticsTaskRunCmd.Flags().String("role", "", "The IAM role that the service assumes to generate statistics.")
		glue_startColumnStatisticsTaskRunCmd.Flags().String("sample-size", "", "The percentage of rows used to generate statistics.")
		glue_startColumnStatisticsTaskRunCmd.Flags().String("security-configuration", "", "Name of the security configuration that is used to encrypt CloudWatch logs for the column stats task run.")
		glue_startColumnStatisticsTaskRunCmd.Flags().String("table-name", "", "The name of the table to generate statistics.")
		glue_startColumnStatisticsTaskRunCmd.MarkFlagRequired("database-name")
		glue_startColumnStatisticsTaskRunCmd.MarkFlagRequired("role")
		glue_startColumnStatisticsTaskRunCmd.MarkFlagRequired("table-name")
	})
	glueCmd.AddCommand(glue_startColumnStatisticsTaskRunCmd)
}
