package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptune_resetDbclusterParameterGroupCmd = &cobra.Command{
	Use:   "reset-dbcluster-parameter-group",
	Short: "Modifies the parameters of a DB cluster parameter group to the default value.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptune_resetDbclusterParameterGroupCmd).Standalone()

	neptune_resetDbclusterParameterGroupCmd.Flags().String("dbcluster-parameter-group-name", "", "The name of the DB cluster parameter group to reset.")
	neptune_resetDbclusterParameterGroupCmd.Flags().Bool("no-reset-all-parameters", false, "A value that is set to `true` to reset all parameters in the DB cluster parameter group to their default values, and `false` otherwise.")
	neptune_resetDbclusterParameterGroupCmd.Flags().String("parameters", "", "A list of parameter names in the DB cluster parameter group to reset to the default values.")
	neptune_resetDbclusterParameterGroupCmd.Flags().Bool("reset-all-parameters", false, "A value that is set to `true` to reset all parameters in the DB cluster parameter group to their default values, and `false` otherwise.")
	neptune_resetDbclusterParameterGroupCmd.MarkFlagRequired("dbcluster-parameter-group-name")
	neptune_resetDbclusterParameterGroupCmd.Flag("no-reset-all-parameters").Hidden = true
	neptuneCmd.AddCommand(neptune_resetDbclusterParameterGroupCmd)
}
