package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_describeDbshardGroupsCmd = &cobra.Command{
	Use:   "describe-dbshard-groups",
	Short: "Describes existing Aurora Limitless Database DB shard groups.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_describeDbshardGroupsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rds_describeDbshardGroupsCmd).Standalone()

		rds_describeDbshardGroupsCmd.Flags().String("dbshard-group-identifier", "", "The user-supplied DB shard group identifier.")
		rds_describeDbshardGroupsCmd.Flags().String("filters", "", "A filter that specifies one or more DB shard groups to describe.")
		rds_describeDbshardGroupsCmd.Flags().String("marker", "", "An optional pagination token provided by a previous `DescribeDBShardGroups` request.")
		rds_describeDbshardGroupsCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
	})
	rdsCmd.AddCommand(rds_describeDbshardGroupsCmd)
}
