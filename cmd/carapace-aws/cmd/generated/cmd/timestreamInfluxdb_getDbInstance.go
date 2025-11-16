package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var timestreamInfluxdb_getDbInstanceCmd = &cobra.Command{
	Use:   "get-db-instance",
	Short: "Returns a Timestream for InfluxDB DB instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(timestreamInfluxdb_getDbInstanceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(timestreamInfluxdb_getDbInstanceCmd).Standalone()

		timestreamInfluxdb_getDbInstanceCmd.Flags().String("identifier", "", "The id of the DB instance.")
		timestreamInfluxdb_getDbInstanceCmd.MarkFlagRequired("identifier")
	})
	timestreamInfluxdbCmd.AddCommand(timestreamInfluxdb_getDbInstanceCmd)
}
