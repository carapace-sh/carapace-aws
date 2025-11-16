package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_resetDbclusterParameterGroupCmd = &cobra.Command{
	Use:   "reset-dbcluster-parameter-group",
	Short: "Modifies the parameters of a DB cluster parameter group to the default value.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_resetDbclusterParameterGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rds_resetDbclusterParameterGroupCmd).Standalone()

		rds_resetDbclusterParameterGroupCmd.Flags().String("dbcluster-parameter-group-name", "", "The name of the DB cluster parameter group to reset.")
		rds_resetDbclusterParameterGroupCmd.Flags().Bool("no-reset-all-parameters", false, "Specifies whether to reset all parameters in the DB cluster parameter group to their default values.")
		rds_resetDbclusterParameterGroupCmd.Flags().String("parameters", "", "A list of parameter names in the DB cluster parameter group to reset to the default values.")
		rds_resetDbclusterParameterGroupCmd.Flags().Bool("reset-all-parameters", false, "Specifies whether to reset all parameters in the DB cluster parameter group to their default values.")
		rds_resetDbclusterParameterGroupCmd.MarkFlagRequired("dbcluster-parameter-group-name")
		rds_resetDbclusterParameterGroupCmd.Flag("no-reset-all-parameters").Hidden = true
	})
	rdsCmd.AddCommand(rds_resetDbclusterParameterGroupCmd)
}
