package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var docdb_createDbclusterParameterGroupCmd = &cobra.Command{
	Use:   "create-dbcluster-parameter-group",
	Short: "Creates a new cluster parameter group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docdb_createDbclusterParameterGroupCmd).Standalone()

	docdb_createDbclusterParameterGroupCmd.Flags().String("dbcluster-parameter-group-name", "", "The name of the cluster parameter group.")
	docdb_createDbclusterParameterGroupCmd.Flags().String("dbparameter-group-family", "", "The cluster parameter group family name.")
	docdb_createDbclusterParameterGroupCmd.Flags().String("description", "", "The description for the cluster parameter group.")
	docdb_createDbclusterParameterGroupCmd.Flags().String("tags", "", "The tags to be assigned to the cluster parameter group.")
	docdb_createDbclusterParameterGroupCmd.MarkFlagRequired("dbcluster-parameter-group-name")
	docdb_createDbclusterParameterGroupCmd.MarkFlagRequired("dbparameter-group-family")
	docdb_createDbclusterParameterGroupCmd.MarkFlagRequired("description")
	docdbCmd.AddCommand(docdb_createDbclusterParameterGroupCmd)
}
