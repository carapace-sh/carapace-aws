package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var docdb_createDbsubnetGroupCmd = &cobra.Command{
	Use:   "create-dbsubnet-group",
	Short: "Creates a new subnet group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docdb_createDbsubnetGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(docdb_createDbsubnetGroupCmd).Standalone()

		docdb_createDbsubnetGroupCmd.Flags().String("dbsubnet-group-description", "", "The description for the subnet group.")
		docdb_createDbsubnetGroupCmd.Flags().String("dbsubnet-group-name", "", "The name for the subnet group.")
		docdb_createDbsubnetGroupCmd.Flags().String("subnet-ids", "", "The Amazon EC2 subnet IDs for the subnet group.")
		docdb_createDbsubnetGroupCmd.Flags().String("tags", "", "The tags to be assigned to the subnet group.")
		docdb_createDbsubnetGroupCmd.MarkFlagRequired("dbsubnet-group-description")
		docdb_createDbsubnetGroupCmd.MarkFlagRequired("dbsubnet-group-name")
		docdb_createDbsubnetGroupCmd.MarkFlagRequired("subnet-ids")
	})
	docdbCmd.AddCommand(docdb_createDbsubnetGroupCmd)
}
