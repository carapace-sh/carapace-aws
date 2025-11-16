package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_deleteDbparameterGroupCmd = &cobra.Command{
	Use:   "delete-dbparameter-group",
	Short: "Deletes a specified DB parameter group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_deleteDbparameterGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rds_deleteDbparameterGroupCmd).Standalone()

		rds_deleteDbparameterGroupCmd.Flags().String("dbparameter-group-name", "", "The name of the DB parameter group.")
		rds_deleteDbparameterGroupCmd.MarkFlagRequired("dbparameter-group-name")
	})
	rdsCmd.AddCommand(rds_deleteDbparameterGroupCmd)
}
