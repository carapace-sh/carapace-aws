package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptune_copyDbclusterParameterGroupCmd = &cobra.Command{
	Use:   "copy-dbcluster-parameter-group",
	Short: "Copies the specified DB cluster parameter group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptune_copyDbclusterParameterGroupCmd).Standalone()

	neptune_copyDbclusterParameterGroupCmd.Flags().String("source-dbcluster-parameter-group-identifier", "", "The identifier or Amazon Resource Name (ARN) for the source DB cluster parameter group.")
	neptune_copyDbclusterParameterGroupCmd.Flags().String("tags", "", "The tags to be assigned to the copied DB cluster parameter group.")
	neptune_copyDbclusterParameterGroupCmd.Flags().String("target-dbcluster-parameter-group-description", "", "A description for the copied DB cluster parameter group.")
	neptune_copyDbclusterParameterGroupCmd.Flags().String("target-dbcluster-parameter-group-identifier", "", "The identifier for the copied DB cluster parameter group.")
	neptune_copyDbclusterParameterGroupCmd.MarkFlagRequired("source-dbcluster-parameter-group-identifier")
	neptune_copyDbclusterParameterGroupCmd.MarkFlagRequired("target-dbcluster-parameter-group-description")
	neptune_copyDbclusterParameterGroupCmd.MarkFlagRequired("target-dbcluster-parameter-group-identifier")
	neptuneCmd.AddCommand(neptune_copyDbclusterParameterGroupCmd)
}
