package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var memorydb_deleteSubnetGroupCmd = &cobra.Command{
	Use:   "delete-subnet-group",
	Short: "Deletes a subnet group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(memorydb_deleteSubnetGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(memorydb_deleteSubnetGroupCmd).Standalone()

		memorydb_deleteSubnetGroupCmd.Flags().String("subnet-group-name", "", "The name of the subnet group to delete.")
		memorydb_deleteSubnetGroupCmd.MarkFlagRequired("subnet-group-name")
	})
	memorydbCmd.AddCommand(memorydb_deleteSubnetGroupCmd)
}
