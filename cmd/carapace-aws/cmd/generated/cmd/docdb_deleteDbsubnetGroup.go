package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var docdb_deleteDbsubnetGroupCmd = &cobra.Command{
	Use:   "delete-dbsubnet-group",
	Short: "Deletes a subnet group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docdb_deleteDbsubnetGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(docdb_deleteDbsubnetGroupCmd).Standalone()

		docdb_deleteDbsubnetGroupCmd.Flags().String("dbsubnet-group-name", "", "The name of the database subnet group to delete.")
		docdb_deleteDbsubnetGroupCmd.MarkFlagRequired("dbsubnet-group-name")
	})
	docdbCmd.AddCommand(docdb_deleteDbsubnetGroupCmd)
}
