package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3tables_getTableMaintenanceConfigurationCmd = &cobra.Command{
	Use:   "get-table-maintenance-configuration",
	Short: "Gets details about the maintenance configuration of a table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3tables_getTableMaintenanceConfigurationCmd).Standalone()

	s3tables_getTableMaintenanceConfigurationCmd.Flags().String("name", "", "The name of the table.")
	s3tables_getTableMaintenanceConfigurationCmd.Flags().String("namespace", "", "The namespace associated with the table.")
	s3tables_getTableMaintenanceConfigurationCmd.Flags().String("table-bucket-arn", "", "The Amazon Resource Name (ARN) of the table bucket.")
	s3tables_getTableMaintenanceConfigurationCmd.MarkFlagRequired("name")
	s3tables_getTableMaintenanceConfigurationCmd.MarkFlagRequired("namespace")
	s3tables_getTableMaintenanceConfigurationCmd.MarkFlagRequired("table-bucket-arn")
	s3tablesCmd.AddCommand(s3tables_getTableMaintenanceConfigurationCmd)
}
