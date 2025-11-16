package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptune_deleteDbclusterParameterGroupCmd = &cobra.Command{
	Use:   "delete-dbcluster-parameter-group",
	Short: "Deletes a specified DB cluster parameter group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptune_deleteDbclusterParameterGroupCmd).Standalone()

	neptune_deleteDbclusterParameterGroupCmd.Flags().String("dbcluster-parameter-group-name", "", "The name of the DB cluster parameter group.")
	neptune_deleteDbclusterParameterGroupCmd.MarkFlagRequired("dbcluster-parameter-group-name")
	neptuneCmd.AddCommand(neptune_deleteDbclusterParameterGroupCmd)
}
