package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var docdb_describeDbclusterParameterGroupsCmd = &cobra.Command{
	Use:   "describe-dbcluster-parameter-groups",
	Short: "Returns a list of `DBClusterParameterGroup` descriptions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docdb_describeDbclusterParameterGroupsCmd).Standalone()

	docdb_describeDbclusterParameterGroupsCmd.Flags().String("dbcluster-parameter-group-name", "", "The name of a specific cluster parameter group to return details for.")
	docdb_describeDbclusterParameterGroupsCmd.Flags().String("filters", "", "This parameter is not currently supported.")
	docdb_describeDbclusterParameterGroupsCmd.Flags().String("marker", "", "An optional pagination token provided by a previous request.")
	docdb_describeDbclusterParameterGroupsCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
	docdbCmd.AddCommand(docdb_describeDbclusterParameterGroupsCmd)
}
