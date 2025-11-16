package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var timestreamInfluxdb_listDbInstancesCmd = &cobra.Command{
	Use:   "list-db-instances",
	Short: "Returns a list of Timestream for InfluxDB DB instances.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(timestreamInfluxdb_listDbInstancesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(timestreamInfluxdb_listDbInstancesCmd).Standalone()

		timestreamInfluxdb_listDbInstancesCmd.Flags().String("max-results", "", "The maximum number of items to return in the output.")
		timestreamInfluxdb_listDbInstancesCmd.Flags().String("next-token", "", "The pagination token.")
	})
	timestreamInfluxdbCmd.AddCommand(timestreamInfluxdb_listDbInstancesCmd)
}
