package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_createDbclusterParameterGroupCmd = &cobra.Command{
	Use:   "create-dbcluster-parameter-group",
	Short: "Creates a new DB cluster parameter group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_createDbclusterParameterGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rds_createDbclusterParameterGroupCmd).Standalone()

		rds_createDbclusterParameterGroupCmd.Flags().String("dbcluster-parameter-group-name", "", "The name of the DB cluster parameter group.")
		rds_createDbclusterParameterGroupCmd.Flags().String("dbparameter-group-family", "", "The DB cluster parameter group family name.")
		rds_createDbclusterParameterGroupCmd.Flags().String("description", "", "The description for the DB cluster parameter group.")
		rds_createDbclusterParameterGroupCmd.Flags().String("tags", "", "Tags to assign to the DB cluster parameter group.")
		rds_createDbclusterParameterGroupCmd.MarkFlagRequired("dbcluster-parameter-group-name")
		rds_createDbclusterParameterGroupCmd.MarkFlagRequired("dbparameter-group-family")
		rds_createDbclusterParameterGroupCmd.MarkFlagRequired("description")
	})
	rdsCmd.AddCommand(rds_createDbclusterParameterGroupCmd)
}
