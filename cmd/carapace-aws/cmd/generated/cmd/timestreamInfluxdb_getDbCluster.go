package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var timestreamInfluxdb_getDbClusterCmd = &cobra.Command{
	Use:   "get-db-cluster",
	Short: "Retrieves information about a Timestream for InfluxDB cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(timestreamInfluxdb_getDbClusterCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(timestreamInfluxdb_getDbClusterCmd).Standalone()

		timestreamInfluxdb_getDbClusterCmd.Flags().String("db-cluster-id", "", "Service-generated unique identifier of the DB cluster to retrieve.")
		timestreamInfluxdb_getDbClusterCmd.MarkFlagRequired("db-cluster-id")
	})
	timestreamInfluxdbCmd.AddCommand(timestreamInfluxdb_getDbClusterCmd)
}
