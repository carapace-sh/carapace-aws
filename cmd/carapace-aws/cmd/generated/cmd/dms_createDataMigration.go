package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_createDataMigrationCmd = &cobra.Command{
	Use:   "create-data-migration",
	Short: "Creates a data migration using the provided settings.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_createDataMigrationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dms_createDataMigrationCmd).Standalone()

		dms_createDataMigrationCmd.Flags().String("data-migration-name", "", "A user-friendly name for the data migration.")
		dms_createDataMigrationCmd.Flags().String("data-migration-type", "", "Specifies if the data migration is full-load only, change data capture (CDC) only, or full-load and CDC.")
		dms_createDataMigrationCmd.Flags().String("enable-cloudwatch-logs", "", "Specifies whether to enable CloudWatch logs for the data migration.")
		dms_createDataMigrationCmd.Flags().String("migration-project-identifier", "", "An identifier for the migration project.")
		dms_createDataMigrationCmd.Flags().String("number-of-jobs", "", "The number of parallel jobs that trigger parallel threads to unload the tables from the source, and then load them to the target.")
		dms_createDataMigrationCmd.Flags().String("selection-rules", "", "An optional JSON string specifying what tables, views, and schemas to include or exclude from the migration.")
		dms_createDataMigrationCmd.Flags().String("service-access-role-arn", "", "The Amazon Resource Name (ARN) for the service access role that you want to use to create the data migration.")
		dms_createDataMigrationCmd.Flags().String("source-data-settings", "", "Specifies information about the source data provider.")
		dms_createDataMigrationCmd.Flags().String("tags", "", "One or more tags to be assigned to the data migration.")
		dms_createDataMigrationCmd.Flags().String("target-data-settings", "", "Specifies information about the target data provider.")
		dms_createDataMigrationCmd.MarkFlagRequired("data-migration-type")
		dms_createDataMigrationCmd.MarkFlagRequired("migration-project-identifier")
		dms_createDataMigrationCmd.MarkFlagRequired("service-access-role-arn")
	})
	dmsCmd.AddCommand(dms_createDataMigrationCmd)
}
