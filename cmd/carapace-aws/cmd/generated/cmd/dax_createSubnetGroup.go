package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dax_createSubnetGroupCmd = &cobra.Command{
	Use:   "create-subnet-group",
	Short: "Creates a new subnet group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dax_createSubnetGroupCmd).Standalone()

	dax_createSubnetGroupCmd.Flags().String("description", "", "A description for the subnet group")
	dax_createSubnetGroupCmd.Flags().String("subnet-group-name", "", "A name for the subnet group.")
	dax_createSubnetGroupCmd.Flags().String("subnet-ids", "", "A list of VPC subnet IDs for the subnet group.")
	dax_createSubnetGroupCmd.MarkFlagRequired("subnet-group-name")
	dax_createSubnetGroupCmd.MarkFlagRequired("subnet-ids")
	daxCmd.AddCommand(dax_createSubnetGroupCmd)
}
