package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_copyDbclusterParameterGroupCmd = &cobra.Command{
	Use:   "copy-dbcluster-parameter-group",
	Short: "Copies the specified DB cluster parameter group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_copyDbclusterParameterGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rds_copyDbclusterParameterGroupCmd).Standalone()

		rds_copyDbclusterParameterGroupCmd.Flags().String("source-dbcluster-parameter-group-identifier", "", "The identifier or Amazon Resource Name (ARN) for the source DB cluster parameter group.")
		rds_copyDbclusterParameterGroupCmd.Flags().String("tags", "", "")
		rds_copyDbclusterParameterGroupCmd.Flags().String("target-dbcluster-parameter-group-description", "", "A description for the copied DB cluster parameter group.")
		rds_copyDbclusterParameterGroupCmd.Flags().String("target-dbcluster-parameter-group-identifier", "", "The identifier for the copied DB cluster parameter group.")
		rds_copyDbclusterParameterGroupCmd.MarkFlagRequired("source-dbcluster-parameter-group-identifier")
		rds_copyDbclusterParameterGroupCmd.MarkFlagRequired("target-dbcluster-parameter-group-description")
		rds_copyDbclusterParameterGroupCmd.MarkFlagRequired("target-dbcluster-parameter-group-identifier")
	})
	rdsCmd.AddCommand(rds_copyDbclusterParameterGroupCmd)
}
