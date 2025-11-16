package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var timestreamInfluxdb_deleteDbInstanceCmd = &cobra.Command{
	Use:   "delete-db-instance",
	Short: "Deletes a Timestream for InfluxDB DB instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(timestreamInfluxdb_deleteDbInstanceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(timestreamInfluxdb_deleteDbInstanceCmd).Standalone()

		timestreamInfluxdb_deleteDbInstanceCmd.Flags().String("identifier", "", "The id of the DB instance.")
		timestreamInfluxdb_deleteDbInstanceCmd.MarkFlagRequired("identifier")
	})
	timestreamInfluxdbCmd.AddCommand(timestreamInfluxdb_deleteDbInstanceCmd)
}
