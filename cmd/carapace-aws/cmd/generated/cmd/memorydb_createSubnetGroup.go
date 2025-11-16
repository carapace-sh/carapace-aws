package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var memorydb_createSubnetGroupCmd = &cobra.Command{
	Use:   "create-subnet-group",
	Short: "Creates a subnet group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(memorydb_createSubnetGroupCmd).Standalone()

	memorydb_createSubnetGroupCmd.Flags().String("description", "", "A description for the subnet group.")
	memorydb_createSubnetGroupCmd.Flags().String("subnet-group-name", "", "The name of the subnet group.")
	memorydb_createSubnetGroupCmd.Flags().String("subnet-ids", "", "A list of VPC subnet IDs for the subnet group.")
	memorydb_createSubnetGroupCmd.Flags().String("tags", "", "A list of tags to be added to this resource.")
	memorydb_createSubnetGroupCmd.MarkFlagRequired("subnet-group-name")
	memorydb_createSubnetGroupCmd.MarkFlagRequired("subnet-ids")
	memorydbCmd.AddCommand(memorydb_createSubnetGroupCmd)
}
