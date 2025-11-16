package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3tables_putTableBucketMaintenanceConfigurationCmd = &cobra.Command{
	Use:   "put-table-bucket-maintenance-configuration",
	Short: "Creates a new maintenance configuration or replaces an existing maintenance configuration for a table bucket.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3tables_putTableBucketMaintenanceConfigurationCmd).Standalone()

	s3tables_putTableBucketMaintenanceConfigurationCmd.Flags().String("table-bucket-arn", "", "The Amazon Resource Name (ARN) of the table bucket associated with the maintenance configuration.")
	s3tables_putTableBucketMaintenanceConfigurationCmd.Flags().String("type", "", "The type of the maintenance configuration.")
	s3tables_putTableBucketMaintenanceConfigurationCmd.Flags().String("value", "", "Defines the values of the maintenance configuration for the table bucket.")
	s3tables_putTableBucketMaintenanceConfigurationCmd.MarkFlagRequired("table-bucket-arn")
	s3tables_putTableBucketMaintenanceConfigurationCmd.MarkFlagRequired("type")
	s3tables_putTableBucketMaintenanceConfigurationCmd.MarkFlagRequired("value")
	s3tablesCmd.AddCommand(s3tables_putTableBucketMaintenanceConfigurationCmd)
}
