package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_describeReplicationTasksCmd = &cobra.Command{
	Use:   "describe-replication-tasks",
	Short: "Returns information about replication tasks for your account in the current region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_describeReplicationTasksCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dms_describeReplicationTasksCmd).Standalone()

		dms_describeReplicationTasksCmd.Flags().String("filters", "", "Filters applied to replication tasks.")
		dms_describeReplicationTasksCmd.Flags().String("marker", "", "An optional pagination token provided by a previous request.")
		dms_describeReplicationTasksCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
		dms_describeReplicationTasksCmd.Flags().String("without-settings", "", "An option to set to avoid returning information about settings.")
	})
	dmsCmd.AddCommand(dms_describeReplicationTasksCmd)
}
