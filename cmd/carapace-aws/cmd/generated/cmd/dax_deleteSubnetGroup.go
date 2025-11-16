package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dax_deleteSubnetGroupCmd = &cobra.Command{
	Use:   "delete-subnet-group",
	Short: "Deletes a subnet group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dax_deleteSubnetGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dax_deleteSubnetGroupCmd).Standalone()

		dax_deleteSubnetGroupCmd.Flags().String("subnet-group-name", "", "The name of the subnet group to delete.")
		dax_deleteSubnetGroupCmd.MarkFlagRequired("subnet-group-name")
	})
	daxCmd.AddCommand(dax_deleteSubnetGroupCmd)
}
