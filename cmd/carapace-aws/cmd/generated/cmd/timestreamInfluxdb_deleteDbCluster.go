package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var timestreamInfluxdb_deleteDbClusterCmd = &cobra.Command{
	Use:   "delete-db-cluster",
	Short: "Deletes a Timestream for InfluxDB cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(timestreamInfluxdb_deleteDbClusterCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(timestreamInfluxdb_deleteDbClusterCmd).Standalone()

		timestreamInfluxdb_deleteDbClusterCmd.Flags().String("db-cluster-id", "", "Service-generated unique identifier of the DB cluster.")
		timestreamInfluxdb_deleteDbClusterCmd.MarkFlagRequired("db-cluster-id")
	})
	timestreamInfluxdbCmd.AddCommand(timestreamInfluxdb_deleteDbClusterCmd)
}
