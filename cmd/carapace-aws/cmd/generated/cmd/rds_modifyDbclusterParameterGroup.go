package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_modifyDbclusterParameterGroupCmd = &cobra.Command{
	Use:   "modify-dbcluster-parameter-group",
	Short: "Modifies the parameters of a DB cluster parameter group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_modifyDbclusterParameterGroupCmd).Standalone()

	rds_modifyDbclusterParameterGroupCmd.Flags().String("dbcluster-parameter-group-name", "", "The name of the DB cluster parameter group to modify.")
	rds_modifyDbclusterParameterGroupCmd.Flags().String("parameters", "", "A list of parameters in the DB cluster parameter group to modify.")
	rds_modifyDbclusterParameterGroupCmd.MarkFlagRequired("dbcluster-parameter-group-name")
	rds_modifyDbclusterParameterGroupCmd.MarkFlagRequired("parameters")
	rdsCmd.AddCommand(rds_modifyDbclusterParameterGroupCmd)
}
