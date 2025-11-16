package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptune_deleteDbparameterGroupCmd = &cobra.Command{
	Use:   "delete-dbparameter-group",
	Short: "Deletes a specified DBParameterGroup.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptune_deleteDbparameterGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(neptune_deleteDbparameterGroupCmd).Standalone()

		neptune_deleteDbparameterGroupCmd.Flags().String("dbparameter-group-name", "", "The name of the DB parameter group.")
		neptune_deleteDbparameterGroupCmd.MarkFlagRequired("dbparameter-group-name")
	})
	neptuneCmd.AddCommand(neptune_deleteDbparameterGroupCmd)
}
