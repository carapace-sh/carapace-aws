package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var timestreamInfluxdb_updateDbClusterCmd = &cobra.Command{
	Use:   "update-db-cluster",
	Short: "Updates a Timestream for InfluxDB cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(timestreamInfluxdb_updateDbClusterCmd).Standalone()

	timestreamInfluxdb_updateDbClusterCmd.Flags().String("db-cluster-id", "", "Service-generated unique identifier of the DB cluster to update.")
	timestreamInfluxdb_updateDbClusterCmd.Flags().String("db-instance-type", "", "Update the DB cluster to use the specified DB instance Type.")
	timestreamInfluxdb_updateDbClusterCmd.Flags().String("db-parameter-group-identifier", "", "Update the DB cluster to use the specified DB parameter group.")
	timestreamInfluxdb_updateDbClusterCmd.Flags().String("failover-mode", "", "Update the DB cluster's failover behavior.")
	timestreamInfluxdb_updateDbClusterCmd.Flags().String("log-delivery-configuration", "", "The log delivery configuration to apply to the DB cluster.")
	timestreamInfluxdb_updateDbClusterCmd.Flags().String("port", "", "Update the DB cluster to use the specified port.")
	timestreamInfluxdb_updateDbClusterCmd.MarkFlagRequired("db-cluster-id")
	timestreamInfluxdbCmd.AddCommand(timestreamInfluxdb_updateDbClusterCmd)
}
