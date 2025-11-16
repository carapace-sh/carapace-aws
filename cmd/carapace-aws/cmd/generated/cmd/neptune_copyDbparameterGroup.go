package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptune_copyDbparameterGroupCmd = &cobra.Command{
	Use:   "copy-dbparameter-group",
	Short: "Copies the specified DB parameter group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptune_copyDbparameterGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(neptune_copyDbparameterGroupCmd).Standalone()

		neptune_copyDbparameterGroupCmd.Flags().String("source-dbparameter-group-identifier", "", "The identifier or ARN for the source DB parameter group.")
		neptune_copyDbparameterGroupCmd.Flags().String("tags", "", "The tags to be assigned to the copied DB parameter group.")
		neptune_copyDbparameterGroupCmd.Flags().String("target-dbparameter-group-description", "", "A description for the copied DB parameter group.")
		neptune_copyDbparameterGroupCmd.Flags().String("target-dbparameter-group-identifier", "", "The identifier for the copied DB parameter group.")
		neptune_copyDbparameterGroupCmd.MarkFlagRequired("source-dbparameter-group-identifier")
		neptune_copyDbparameterGroupCmd.MarkFlagRequired("target-dbparameter-group-description")
		neptune_copyDbparameterGroupCmd.MarkFlagRequired("target-dbparameter-group-identifier")
	})
	neptuneCmd.AddCommand(neptune_copyDbparameterGroupCmd)
}
