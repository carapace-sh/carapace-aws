package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var timestreamInfluxdb_getDbParameterGroupCmd = &cobra.Command{
	Use:   "get-db-parameter-group",
	Short: "Returns a Timestream for InfluxDB DB parameter group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(timestreamInfluxdb_getDbParameterGroupCmd).Standalone()

	timestreamInfluxdb_getDbParameterGroupCmd.Flags().String("identifier", "", "The id of the DB parameter group.")
	timestreamInfluxdb_getDbParameterGroupCmd.MarkFlagRequired("identifier")
	timestreamInfluxdbCmd.AddCommand(timestreamInfluxdb_getDbParameterGroupCmd)
}
