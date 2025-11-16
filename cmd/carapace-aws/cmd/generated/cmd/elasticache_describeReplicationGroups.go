package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticache_describeReplicationGroupsCmd = &cobra.Command{
	Use:   "describe-replication-groups",
	Short: "Returns information about a particular replication group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticache_describeReplicationGroupsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elasticache_describeReplicationGroupsCmd).Standalone()

		elasticache_describeReplicationGroupsCmd.Flags().String("marker", "", "An optional marker returned from a prior request.")
		elasticache_describeReplicationGroupsCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
		elasticache_describeReplicationGroupsCmd.Flags().String("replication-group-id", "", "The identifier for the replication group to be described.")
	})
	elasticacheCmd.AddCommand(elasticache_describeReplicationGroupsCmd)
}
