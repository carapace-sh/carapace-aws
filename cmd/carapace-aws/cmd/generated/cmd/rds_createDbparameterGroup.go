package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_createDbparameterGroupCmd = &cobra.Command{
	Use:   "create-dbparameter-group",
	Short: "Creates a new DB parameter group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_createDbparameterGroupCmd).Standalone()

	rds_createDbparameterGroupCmd.Flags().String("dbparameter-group-family", "", "The DB parameter group family name.")
	rds_createDbparameterGroupCmd.Flags().String("dbparameter-group-name", "", "The name of the DB parameter group.")
	rds_createDbparameterGroupCmd.Flags().String("description", "", "The description for the DB parameter group.")
	rds_createDbparameterGroupCmd.Flags().String("tags", "", "Tags to assign to the DB parameter group.")
	rds_createDbparameterGroupCmd.MarkFlagRequired("dbparameter-group-family")
	rds_createDbparameterGroupCmd.MarkFlagRequired("dbparameter-group-name")
	rds_createDbparameterGroupCmd.MarkFlagRequired("description")
	rdsCmd.AddCommand(rds_createDbparameterGroupCmd)
}
