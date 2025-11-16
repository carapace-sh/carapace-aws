package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_resetDbparameterGroupCmd = &cobra.Command{
	Use:   "reset-dbparameter-group",
	Short: "Modifies the parameters of a DB parameter group to the engine/system default value.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_resetDbparameterGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rds_resetDbparameterGroupCmd).Standalone()

		rds_resetDbparameterGroupCmd.Flags().String("dbparameter-group-name", "", "The name of the DB parameter group.")
		rds_resetDbparameterGroupCmd.Flags().Bool("no-reset-all-parameters", false, "Specifies whether to reset all parameters in the DB parameter group to default values.")
		rds_resetDbparameterGroupCmd.Flags().String("parameters", "", "To reset the entire DB parameter group, specify the `DBParameterGroup` name and `ResetAllParameters` parameters.")
		rds_resetDbparameterGroupCmd.Flags().Bool("reset-all-parameters", false, "Specifies whether to reset all parameters in the DB parameter group to default values.")
		rds_resetDbparameterGroupCmd.MarkFlagRequired("dbparameter-group-name")
		rds_resetDbparameterGroupCmd.Flag("no-reset-all-parameters").Hidden = true
	})
	rdsCmd.AddCommand(rds_resetDbparameterGroupCmd)
}
