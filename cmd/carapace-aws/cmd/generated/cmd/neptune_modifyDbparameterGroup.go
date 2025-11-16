package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptune_modifyDbparameterGroupCmd = &cobra.Command{
	Use:   "modify-dbparameter-group",
	Short: "Modifies the parameters of a DB parameter group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptune_modifyDbparameterGroupCmd).Standalone()

	neptune_modifyDbparameterGroupCmd.Flags().String("dbparameter-group-name", "", "The name of the DB parameter group.")
	neptune_modifyDbparameterGroupCmd.Flags().String("parameters", "", "An array of parameter names, values, and the apply method for the parameter update.")
	neptune_modifyDbparameterGroupCmd.MarkFlagRequired("dbparameter-group-name")
	neptune_modifyDbparameterGroupCmd.MarkFlagRequired("parameters")
	neptuneCmd.AddCommand(neptune_modifyDbparameterGroupCmd)
}
