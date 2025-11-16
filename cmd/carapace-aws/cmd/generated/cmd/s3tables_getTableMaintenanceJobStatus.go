package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3tables_getTableMaintenanceJobStatusCmd = &cobra.Command{
	Use:   "get-table-maintenance-job-status",
	Short: "Gets the status of a maintenance job for a table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3tables_getTableMaintenanceJobStatusCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3tables_getTableMaintenanceJobStatusCmd).Standalone()

		s3tables_getTableMaintenanceJobStatusCmd.Flags().String("name", "", "The name of the table containing the maintenance job status you want to check.")
		s3tables_getTableMaintenanceJobStatusCmd.Flags().String("namespace", "", "The name of the namespace the table is associated with.")
		s3tables_getTableMaintenanceJobStatusCmd.Flags().String("table-bucket-arn", "", "The Amazon Resource Name (ARN) of the table bucket.")
		s3tables_getTableMaintenanceJobStatusCmd.MarkFlagRequired("name")
		s3tables_getTableMaintenanceJobStatusCmd.MarkFlagRequired("namespace")
		s3tables_getTableMaintenanceJobStatusCmd.MarkFlagRequired("table-bucket-arn")
	})
	s3tablesCmd.AddCommand(s3tables_getTableMaintenanceJobStatusCmd)
}
