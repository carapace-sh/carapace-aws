package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_describeReplicationSubnetGroupsCmd = &cobra.Command{
	Use:   "describe-replication-subnet-groups",
	Short: "Returns information about the replication subnet groups.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_describeReplicationSubnetGroupsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dms_describeReplicationSubnetGroupsCmd).Standalone()

		dms_describeReplicationSubnetGroupsCmd.Flags().String("filters", "", "Filters applied to replication subnet groups.")
		dms_describeReplicationSubnetGroupsCmd.Flags().String("marker", "", "An optional pagination token provided by a previous request.")
		dms_describeReplicationSubnetGroupsCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
	})
	dmsCmd.AddCommand(dms_describeReplicationSubnetGroupsCmd)
}
