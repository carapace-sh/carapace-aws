package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var timestreamInfluxdb_updateDbInstanceCmd = &cobra.Command{
	Use:   "update-db-instance",
	Short: "Updates a Timestream for InfluxDB DB instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(timestreamInfluxdb_updateDbInstanceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(timestreamInfluxdb_updateDbInstanceCmd).Standalone()

		timestreamInfluxdb_updateDbInstanceCmd.Flags().String("allocated-storage", "", "The amount of storage to allocate for your DB storage type (in gibibytes).")
		timestreamInfluxdb_updateDbInstanceCmd.Flags().String("db-instance-type", "", "The Timestream for InfluxDB DB instance type to run InfluxDB on.")
		timestreamInfluxdb_updateDbInstanceCmd.Flags().String("db-parameter-group-identifier", "", "The id of the DB parameter group to assign to your DB instance.")
		timestreamInfluxdb_updateDbInstanceCmd.Flags().String("db-storage-type", "", "The Timestream for InfluxDB DB storage type that InfluxDB stores data on.")
		timestreamInfluxdb_updateDbInstanceCmd.Flags().String("deployment-type", "", "Specifies whether the DB instance will be deployed as a standalone instance or with a Multi-AZ standby for high availability.")
		timestreamInfluxdb_updateDbInstanceCmd.Flags().String("identifier", "", "The id of the DB instance.")
		timestreamInfluxdb_updateDbInstanceCmd.Flags().String("log-delivery-configuration", "", "Configuration for sending InfluxDB engine logs to send to specified S3 bucket.")
		timestreamInfluxdb_updateDbInstanceCmd.Flags().String("port", "", "The port number on which InfluxDB accepts connections.")
		timestreamInfluxdb_updateDbInstanceCmd.MarkFlagRequired("identifier")
	})
	timestreamInfluxdbCmd.AddCommand(timestreamInfluxdb_updateDbInstanceCmd)
}
