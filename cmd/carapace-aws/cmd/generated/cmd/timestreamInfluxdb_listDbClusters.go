package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var timestreamInfluxdb_listDbClustersCmd = &cobra.Command{
	Use:   "list-db-clusters",
	Short: "Returns a list of Timestream for InfluxDB DB clusters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(timestreamInfluxdb_listDbClustersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(timestreamInfluxdb_listDbClustersCmd).Standalone()

		timestreamInfluxdb_listDbClustersCmd.Flags().String("max-results", "", "The maximum number of items to return in the output.")
		timestreamInfluxdb_listDbClustersCmd.Flags().String("next-token", "", "The pagination token.")
	})
	timestreamInfluxdbCmd.AddCommand(timestreamInfluxdb_listDbClustersCmd)
}
