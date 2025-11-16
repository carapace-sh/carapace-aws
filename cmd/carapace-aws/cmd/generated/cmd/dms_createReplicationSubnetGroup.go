package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_createReplicationSubnetGroupCmd = &cobra.Command{
	Use:   "create-replication-subnet-group",
	Short: "Creates a replication subnet group given a list of the subnet IDs in a VPC.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_createReplicationSubnetGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dms_createReplicationSubnetGroupCmd).Standalone()

		dms_createReplicationSubnetGroupCmd.Flags().String("replication-subnet-group-description", "", "The description for the subnet group.")
		dms_createReplicationSubnetGroupCmd.Flags().String("replication-subnet-group-identifier", "", "The name for the replication subnet group.")
		dms_createReplicationSubnetGroupCmd.Flags().String("subnet-ids", "", "Two or more subnet IDs to be assigned to the subnet group.")
		dms_createReplicationSubnetGroupCmd.Flags().String("tags", "", "One or more tags to be assigned to the subnet group.")
		dms_createReplicationSubnetGroupCmd.MarkFlagRequired("replication-subnet-group-description")
		dms_createReplicationSubnetGroupCmd.MarkFlagRequired("replication-subnet-group-identifier")
		dms_createReplicationSubnetGroupCmd.MarkFlagRequired("subnet-ids")
	})
	dmsCmd.AddCommand(dms_createReplicationSubnetGroupCmd)
}
