package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_deleteReplicationSubnetGroupCmd = &cobra.Command{
	Use:   "delete-replication-subnet-group",
	Short: "Deletes a subnet group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_deleteReplicationSubnetGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dms_deleteReplicationSubnetGroupCmd).Standalone()

		dms_deleteReplicationSubnetGroupCmd.Flags().String("replication-subnet-group-identifier", "", "The subnet group name of the replication instance.")
		dms_deleteReplicationSubnetGroupCmd.MarkFlagRequired("replication-subnet-group-identifier")
	})
	dmsCmd.AddCommand(dms_deleteReplicationSubnetGroupCmd)
}
