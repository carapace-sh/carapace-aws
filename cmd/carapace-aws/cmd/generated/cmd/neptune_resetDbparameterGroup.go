package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptune_resetDbparameterGroupCmd = &cobra.Command{
	Use:   "reset-dbparameter-group",
	Short: "Modifies the parameters of a DB parameter group to the engine/system default value.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptune_resetDbparameterGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(neptune_resetDbparameterGroupCmd).Standalone()

		neptune_resetDbparameterGroupCmd.Flags().String("dbparameter-group-name", "", "The name of the DB parameter group.")
		neptune_resetDbparameterGroupCmd.Flags().Bool("no-reset-all-parameters", false, "Specifies whether (`true`) or not (`false`) to reset all parameters in the DB parameter group to default values.")
		neptune_resetDbparameterGroupCmd.Flags().String("parameters", "", "To reset the entire DB parameter group, specify the `DBParameterGroup` name and `ResetAllParameters` parameters.")
		neptune_resetDbparameterGroupCmd.Flags().Bool("reset-all-parameters", false, "Specifies whether (`true`) or not (`false`) to reset all parameters in the DB parameter group to default values.")
		neptune_resetDbparameterGroupCmd.MarkFlagRequired("dbparameter-group-name")
		neptune_resetDbparameterGroupCmd.Flag("no-reset-all-parameters").Hidden = true
	})
	neptuneCmd.AddCommand(neptune_resetDbparameterGroupCmd)
}
