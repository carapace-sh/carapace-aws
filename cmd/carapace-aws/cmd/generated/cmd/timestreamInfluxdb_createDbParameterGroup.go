package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var timestreamInfluxdb_createDbParameterGroupCmd = &cobra.Command{
	Use:   "create-db-parameter-group",
	Short: "Creates a new Timestream for InfluxDB DB parameter group to associate with DB instances.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(timestreamInfluxdb_createDbParameterGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(timestreamInfluxdb_createDbParameterGroupCmd).Standalone()

		timestreamInfluxdb_createDbParameterGroupCmd.Flags().String("description", "", "A description of the DB parameter group.")
		timestreamInfluxdb_createDbParameterGroupCmd.Flags().String("name", "", "The name of the DB parameter group.")
		timestreamInfluxdb_createDbParameterGroupCmd.Flags().String("parameters", "", "A list of the parameters that comprise the DB parameter group.")
		timestreamInfluxdb_createDbParameterGroupCmd.Flags().String("tags", "", "A list of key-value pairs to associate with the DB parameter group.")
		timestreamInfluxdb_createDbParameterGroupCmd.MarkFlagRequired("name")
	})
	timestreamInfluxdbCmd.AddCommand(timestreamInfluxdb_createDbParameterGroupCmd)
}
