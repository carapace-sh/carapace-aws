package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dax_updateSubnetGroupCmd = &cobra.Command{
	Use:   "update-subnet-group",
	Short: "Modifies an existing subnet group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dax_updateSubnetGroupCmd).Standalone()

	dax_updateSubnetGroupCmd.Flags().String("description", "", "A description of the subnet group.")
	dax_updateSubnetGroupCmd.Flags().String("subnet-group-name", "", "The name of the subnet group.")
	dax_updateSubnetGroupCmd.Flags().String("subnet-ids", "", "A list of subnet IDs in the subnet group.")
	dax_updateSubnetGroupCmd.MarkFlagRequired("subnet-group-name")
	daxCmd.AddCommand(dax_updateSubnetGroupCmd)
}
