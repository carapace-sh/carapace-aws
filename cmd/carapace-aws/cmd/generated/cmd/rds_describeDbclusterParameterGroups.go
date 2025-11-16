package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_describeDbclusterParameterGroupsCmd = &cobra.Command{
	Use:   "describe-dbcluster-parameter-groups",
	Short: "Returns a list of `DBClusterParameterGroup` descriptions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_describeDbclusterParameterGroupsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rds_describeDbclusterParameterGroupsCmd).Standalone()

		rds_describeDbclusterParameterGroupsCmd.Flags().String("dbcluster-parameter-group-name", "", "The name of a specific DB cluster parameter group to return details for.")
		rds_describeDbclusterParameterGroupsCmd.Flags().String("filters", "", "This parameter isn't currently supported.")
		rds_describeDbclusterParameterGroupsCmd.Flags().String("marker", "", "An optional pagination token provided by a previous `DescribeDBClusterParameterGroups` request.")
		rds_describeDbclusterParameterGroupsCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
	})
	rdsCmd.AddCommand(rds_describeDbclusterParameterGroupsCmd)
}
