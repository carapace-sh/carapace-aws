package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var timestreamInfluxdb_listDbParameterGroupsCmd = &cobra.Command{
	Use:   "list-db-parameter-groups",
	Short: "Returns a list of Timestream for InfluxDB DB parameter groups.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(timestreamInfluxdb_listDbParameterGroupsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(timestreamInfluxdb_listDbParameterGroupsCmd).Standalone()

		timestreamInfluxdb_listDbParameterGroupsCmd.Flags().String("max-results", "", "The maximum number of items to return in the output.")
		timestreamInfluxdb_listDbParameterGroupsCmd.Flags().String("next-token", "", "The pagination token.")
	})
	timestreamInfluxdbCmd.AddCommand(timestreamInfluxdb_listDbParameterGroupsCmd)
}
