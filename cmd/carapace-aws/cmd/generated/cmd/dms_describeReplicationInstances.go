package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_describeReplicationInstancesCmd = &cobra.Command{
	Use:   "describe-replication-instances",
	Short: "Returns information about replication instances for your account in the current region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_describeReplicationInstancesCmd).Standalone()

	dms_describeReplicationInstancesCmd.Flags().String("filters", "", "Filters applied to replication instances.")
	dms_describeReplicationInstancesCmd.Flags().String("marker", "", "An optional pagination token provided by a previous request.")
	dms_describeReplicationInstancesCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
	dmsCmd.AddCommand(dms_describeReplicationInstancesCmd)
}
