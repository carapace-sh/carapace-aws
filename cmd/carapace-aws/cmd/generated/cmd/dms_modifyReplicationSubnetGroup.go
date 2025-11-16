package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_modifyReplicationSubnetGroupCmd = &cobra.Command{
	Use:   "modify-replication-subnet-group",
	Short: "Modifies the settings for the specified replication subnet group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_modifyReplicationSubnetGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dms_modifyReplicationSubnetGroupCmd).Standalone()

		dms_modifyReplicationSubnetGroupCmd.Flags().String("replication-subnet-group-description", "", "A description for the replication instance subnet group.")
		dms_modifyReplicationSubnetGroupCmd.Flags().String("replication-subnet-group-identifier", "", "The name of the replication instance subnet group.")
		dms_modifyReplicationSubnetGroupCmd.Flags().String("subnet-ids", "", "A list of subnet IDs.")
		dms_modifyReplicationSubnetGroupCmd.MarkFlagRequired("replication-subnet-group-identifier")
		dms_modifyReplicationSubnetGroupCmd.MarkFlagRequired("subnet-ids")
	})
	dmsCmd.AddCommand(dms_modifyReplicationSubnetGroupCmd)
}
