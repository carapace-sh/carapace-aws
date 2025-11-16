package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptune_describeDbclusterParameterGroupsCmd = &cobra.Command{
	Use:   "describe-dbcluster-parameter-groups",
	Short: "Returns a list of `DBClusterParameterGroup` descriptions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptune_describeDbclusterParameterGroupsCmd).Standalone()

	neptune_describeDbclusterParameterGroupsCmd.Flags().String("dbcluster-parameter-group-name", "", "The name of a specific DB cluster parameter group to return details for.")
	neptune_describeDbclusterParameterGroupsCmd.Flags().String("filters", "", "This parameter is not currently supported.")
	neptune_describeDbclusterParameterGroupsCmd.Flags().String("marker", "", "An optional pagination token provided by a previous `DescribeDBClusterParameterGroups` request.")
	neptune_describeDbclusterParameterGroupsCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
	neptuneCmd.AddCommand(neptune_describeDbclusterParameterGroupsCmd)
}
