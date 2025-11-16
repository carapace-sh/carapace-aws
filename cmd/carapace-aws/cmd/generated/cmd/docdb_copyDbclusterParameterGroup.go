package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var docdb_copyDbclusterParameterGroupCmd = &cobra.Command{
	Use:   "copy-dbcluster-parameter-group",
	Short: "Copies the specified cluster parameter group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docdb_copyDbclusterParameterGroupCmd).Standalone()

	docdb_copyDbclusterParameterGroupCmd.Flags().String("source-dbcluster-parameter-group-identifier", "", "The identifier or Amazon Resource Name (ARN) for the source cluster parameter group.")
	docdb_copyDbclusterParameterGroupCmd.Flags().String("tags", "", "The tags that are to be assigned to the parameter group.")
	docdb_copyDbclusterParameterGroupCmd.Flags().String("target-dbcluster-parameter-group-description", "", "A description for the copied cluster parameter group.")
	docdb_copyDbclusterParameterGroupCmd.Flags().String("target-dbcluster-parameter-group-identifier", "", "The identifier for the copied cluster parameter group.")
	docdb_copyDbclusterParameterGroupCmd.MarkFlagRequired("source-dbcluster-parameter-group-identifier")
	docdb_copyDbclusterParameterGroupCmd.MarkFlagRequired("target-dbcluster-parameter-group-description")
	docdb_copyDbclusterParameterGroupCmd.MarkFlagRequired("target-dbcluster-parameter-group-identifier")
	docdbCmd.AddCommand(docdb_copyDbclusterParameterGroupCmd)
}
