package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var docdb_deleteDbclusterParameterGroupCmd = &cobra.Command{
	Use:   "delete-dbcluster-parameter-group",
	Short: "Deletes a specified cluster parameter group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docdb_deleteDbclusterParameterGroupCmd).Standalone()

	docdb_deleteDbclusterParameterGroupCmd.Flags().String("dbcluster-parameter-group-name", "", "The name of the cluster parameter group.")
	docdb_deleteDbclusterParameterGroupCmd.MarkFlagRequired("dbcluster-parameter-group-name")
	docdbCmd.AddCommand(docdb_deleteDbclusterParameterGroupCmd)
}
