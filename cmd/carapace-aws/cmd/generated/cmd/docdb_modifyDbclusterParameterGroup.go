package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var docdb_modifyDbclusterParameterGroupCmd = &cobra.Command{
	Use:   "modify-dbcluster-parameter-group",
	Short: "Modifies the parameters of a cluster parameter group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docdb_modifyDbclusterParameterGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(docdb_modifyDbclusterParameterGroupCmd).Standalone()

		docdb_modifyDbclusterParameterGroupCmd.Flags().String("dbcluster-parameter-group-name", "", "The name of the cluster parameter group to modify.")
		docdb_modifyDbclusterParameterGroupCmd.Flags().String("parameters", "", "A list of parameters in the cluster parameter group to modify.")
		docdb_modifyDbclusterParameterGroupCmd.MarkFlagRequired("dbcluster-parameter-group-name")
		docdb_modifyDbclusterParameterGroupCmd.MarkFlagRequired("parameters")
	})
	docdbCmd.AddCommand(docdb_modifyDbclusterParameterGroupCmd)
}
