package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_copyDbparameterGroupCmd = &cobra.Command{
	Use:   "copy-dbparameter-group",
	Short: "Copies the specified DB parameter group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_copyDbparameterGroupCmd).Standalone()

	rds_copyDbparameterGroupCmd.Flags().String("source-dbparameter-group-identifier", "", "The identifier or ARN for the source DB parameter group.")
	rds_copyDbparameterGroupCmd.Flags().String("tags", "", "")
	rds_copyDbparameterGroupCmd.Flags().String("target-dbparameter-group-description", "", "A description for the copied DB parameter group.")
	rds_copyDbparameterGroupCmd.Flags().String("target-dbparameter-group-identifier", "", "The identifier for the copied DB parameter group.")
	rds_copyDbparameterGroupCmd.MarkFlagRequired("source-dbparameter-group-identifier")
	rds_copyDbparameterGroupCmd.MarkFlagRequired("target-dbparameter-group-description")
	rds_copyDbparameterGroupCmd.MarkFlagRequired("target-dbparameter-group-identifier")
	rdsCmd.AddCommand(rds_copyDbparameterGroupCmd)
}
