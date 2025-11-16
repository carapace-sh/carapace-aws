package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_createColumnStatisticsTaskSettingsCmd = &cobra.Command{
	Use:   "create-column-statistics-task-settings",
	Short: "Creates settings for a column statistics task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_createColumnStatisticsTaskSettingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_createColumnStatisticsTaskSettingsCmd).Standalone()

		glue_createColumnStatisticsTaskSettingsCmd.Flags().String("catalog-id", "", "The ID of the Data Catalog in which the database resides.")
		glue_createColumnStatisticsTaskSettingsCmd.Flags().String("column-name-list", "", "A list of column names for which to run statistics.")
		glue_createColumnStatisticsTaskSettingsCmd.Flags().String("database-name", "", "The name of the database where the table resides.")
		glue_createColumnStatisticsTaskSettingsCmd.Flags().String("role", "", "The role used for running the column statistics.")
		glue_createColumnStatisticsTaskSettingsCmd.Flags().String("sample-size", "", "The percentage of data to sample.")
		glue_createColumnStatisticsTaskSettingsCmd.Flags().String("schedule", "", "A schedule for running the column statistics, specified in CRON syntax.")
		glue_createColumnStatisticsTaskSettingsCmd.Flags().String("security-configuration", "", "Name of the security configuration that is used to encrypt CloudWatch logs.")
		glue_createColumnStatisticsTaskSettingsCmd.Flags().String("table-name", "", "The name of the table for which to generate column statistics.")
		glue_createColumnStatisticsTaskSettingsCmd.Flags().String("tags", "", "A map of tags.")
		glue_createColumnStatisticsTaskSettingsCmd.MarkFlagRequired("database-name")
		glue_createColumnStatisticsTaskSettingsCmd.MarkFlagRequired("role")
		glue_createColumnStatisticsTaskSettingsCmd.MarkFlagRequired("table-name")
	})
	glueCmd.AddCommand(glue_createColumnStatisticsTaskSettingsCmd)
}
