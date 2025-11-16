package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_modifyDbparameterGroupCmd = &cobra.Command{
	Use:   "modify-dbparameter-group",
	Short: "Modifies the parameters of a DB parameter group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_modifyDbparameterGroupCmd).Standalone()

	rds_modifyDbparameterGroupCmd.Flags().String("dbparameter-group-name", "", "The name of the DB parameter group.")
	rds_modifyDbparameterGroupCmd.Flags().String("parameters", "", "An array of parameter names, values, and the application methods for the parameter update.")
	rds_modifyDbparameterGroupCmd.MarkFlagRequired("dbparameter-group-name")
	rds_modifyDbparameterGroupCmd.MarkFlagRequired("parameters")
	rdsCmd.AddCommand(rds_modifyDbparameterGroupCmd)
}
