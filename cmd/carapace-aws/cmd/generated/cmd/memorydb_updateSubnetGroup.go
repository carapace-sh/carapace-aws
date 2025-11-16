package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var memorydb_updateSubnetGroupCmd = &cobra.Command{
	Use:   "update-subnet-group",
	Short: "Updates a subnet group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(memorydb_updateSubnetGroupCmd).Standalone()

	memorydb_updateSubnetGroupCmd.Flags().String("description", "", "A description of the subnet group")
	memorydb_updateSubnetGroupCmd.Flags().String("subnet-group-name", "", "The name of the subnet group")
	memorydb_updateSubnetGroupCmd.Flags().String("subnet-ids", "", "The EC2 subnet IDs for the subnet group.")
	memorydb_updateSubnetGroupCmd.MarkFlagRequired("subnet-group-name")
	memorydbCmd.AddCommand(memorydb_updateSubnetGroupCmd)
}
