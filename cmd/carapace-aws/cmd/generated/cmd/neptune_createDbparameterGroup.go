package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptune_createDbparameterGroupCmd = &cobra.Command{
	Use:   "create-dbparameter-group",
	Short: "Creates a new DB parameter group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptune_createDbparameterGroupCmd).Standalone()

	neptune_createDbparameterGroupCmd.Flags().String("dbparameter-group-family", "", "The DB parameter group family name.")
	neptune_createDbparameterGroupCmd.Flags().String("dbparameter-group-name", "", "The name of the DB parameter group.")
	neptune_createDbparameterGroupCmd.Flags().String("description", "", "The description for the DB parameter group.")
	neptune_createDbparameterGroupCmd.Flags().String("tags", "", "The tags to be assigned to the new DB parameter group.")
	neptune_createDbparameterGroupCmd.MarkFlagRequired("dbparameter-group-family")
	neptune_createDbparameterGroupCmd.MarkFlagRequired("dbparameter-group-name")
	neptune_createDbparameterGroupCmd.MarkFlagRequired("description")
	neptuneCmd.AddCommand(neptune_createDbparameterGroupCmd)
}
