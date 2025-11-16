package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var docdb_modifyDbsubnetGroupCmd = &cobra.Command{
	Use:   "modify-dbsubnet-group",
	Short: "Modifies an existing subnet group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docdb_modifyDbsubnetGroupCmd).Standalone()

	docdb_modifyDbsubnetGroupCmd.Flags().String("dbsubnet-group-description", "", "The description for the subnet group.")
	docdb_modifyDbsubnetGroupCmd.Flags().String("dbsubnet-group-name", "", "The name for the subnet group.")
	docdb_modifyDbsubnetGroupCmd.Flags().String("subnet-ids", "", "The Amazon EC2 subnet IDs for the subnet group.")
	docdb_modifyDbsubnetGroupCmd.MarkFlagRequired("dbsubnet-group-name")
	docdb_modifyDbsubnetGroupCmd.MarkFlagRequired("subnet-ids")
	docdbCmd.AddCommand(docdb_modifyDbsubnetGroupCmd)
}
