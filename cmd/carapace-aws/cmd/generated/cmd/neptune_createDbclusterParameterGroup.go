package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptune_createDbclusterParameterGroupCmd = &cobra.Command{
	Use:   "create-dbcluster-parameter-group",
	Short: "Creates a new DB cluster parameter group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptune_createDbclusterParameterGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(neptune_createDbclusterParameterGroupCmd).Standalone()

		neptune_createDbclusterParameterGroupCmd.Flags().String("dbcluster-parameter-group-name", "", "The name of the DB cluster parameter group.")
		neptune_createDbclusterParameterGroupCmd.Flags().String("dbparameter-group-family", "", "The DB cluster parameter group family name.")
		neptune_createDbclusterParameterGroupCmd.Flags().String("description", "", "The description for the DB cluster parameter group.")
		neptune_createDbclusterParameterGroupCmd.Flags().String("tags", "", "The tags to be assigned to the new DB cluster parameter group.")
		neptune_createDbclusterParameterGroupCmd.MarkFlagRequired("dbcluster-parameter-group-name")
		neptune_createDbclusterParameterGroupCmd.MarkFlagRequired("dbparameter-group-family")
		neptune_createDbclusterParameterGroupCmd.MarkFlagRequired("description")
	})
	neptuneCmd.AddCommand(neptune_createDbclusterParameterGroupCmd)
}
