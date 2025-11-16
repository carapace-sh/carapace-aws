package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_deleteDbclusterParameterGroupCmd = &cobra.Command{
	Use:   "delete-dbcluster-parameter-group",
	Short: "Deletes a specified DB cluster parameter group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_deleteDbclusterParameterGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rds_deleteDbclusterParameterGroupCmd).Standalone()

		rds_deleteDbclusterParameterGroupCmd.Flags().String("dbcluster-parameter-group-name", "", "The name of the DB cluster parameter group.")
		rds_deleteDbclusterParameterGroupCmd.MarkFlagRequired("dbcluster-parameter-group-name")
	})
	rdsCmd.AddCommand(rds_deleteDbclusterParameterGroupCmd)
}
