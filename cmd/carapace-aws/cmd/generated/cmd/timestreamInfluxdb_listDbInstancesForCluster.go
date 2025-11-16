package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var timestreamInfluxdb_listDbInstancesForClusterCmd = &cobra.Command{
	Use:   "list-db-instances-for-cluster",
	Short: "Returns a list of Timestream for InfluxDB clusters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(timestreamInfluxdb_listDbInstancesForClusterCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(timestreamInfluxdb_listDbInstancesForClusterCmd).Standalone()

		timestreamInfluxdb_listDbInstancesForClusterCmd.Flags().String("db-cluster-id", "", "Service-generated unique identifier of the DB cluster.")
		timestreamInfluxdb_listDbInstancesForClusterCmd.Flags().String("max-results", "", "The maximum number of items to return in the output.")
		timestreamInfluxdb_listDbInstancesForClusterCmd.Flags().String("next-token", "", "The pagination token.")
		timestreamInfluxdb_listDbInstancesForClusterCmd.MarkFlagRequired("db-cluster-id")
	})
	timestreamInfluxdbCmd.AddCommand(timestreamInfluxdb_listDbInstancesForClusterCmd)
}
