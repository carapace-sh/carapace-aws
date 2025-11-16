package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptune_modifyDbclusterParameterGroupCmd = &cobra.Command{
	Use:   "modify-dbcluster-parameter-group",
	Short: "Modifies the parameters of a DB cluster parameter group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptune_modifyDbclusterParameterGroupCmd).Standalone()

	neptune_modifyDbclusterParameterGroupCmd.Flags().String("dbcluster-parameter-group-name", "", "The name of the DB cluster parameter group to modify.")
	neptune_modifyDbclusterParameterGroupCmd.Flags().String("parameters", "", "A list of parameters in the DB cluster parameter group to modify.")
	neptune_modifyDbclusterParameterGroupCmd.MarkFlagRequired("dbcluster-parameter-group-name")
	neptune_modifyDbclusterParameterGroupCmd.MarkFlagRequired("parameters")
	neptuneCmd.AddCommand(neptune_modifyDbclusterParameterGroupCmd)
}
