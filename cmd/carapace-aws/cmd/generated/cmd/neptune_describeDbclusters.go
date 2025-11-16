package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptune_describeDbclustersCmd = &cobra.Command{
	Use:   "describe-dbclusters",
	Short: "Returns information about provisioned DB clusters, and supports pagination.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptune_describeDbclustersCmd).Standalone()

	neptune_describeDbclustersCmd.Flags().String("dbcluster-identifier", "", "The user-supplied DB cluster identifier.")
	neptune_describeDbclustersCmd.Flags().String("filters", "", "A filter that specifies one or more DB clusters to describe.")
	neptune_describeDbclustersCmd.Flags().String("marker", "", "An optional pagination token provided by a previous [DescribeDBClusters]() request.")
	neptune_describeDbclustersCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
	neptuneCmd.AddCommand(neptune_describeDbclustersCmd)
}
