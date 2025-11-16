package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticache_describeGlobalReplicationGroupsCmd = &cobra.Command{
	Use:   "describe-global-replication-groups",
	Short: "Returns information about a particular global replication group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticache_describeGlobalReplicationGroupsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elasticache_describeGlobalReplicationGroupsCmd).Standalone()

		elasticache_describeGlobalReplicationGroupsCmd.Flags().String("global-replication-group-id", "", "The name of the Global datastore")
		elasticache_describeGlobalReplicationGroupsCmd.Flags().String("marker", "", "An optional marker returned from a prior request.")
		elasticache_describeGlobalReplicationGroupsCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
		elasticache_describeGlobalReplicationGroupsCmd.Flags().String("show-member-info", "", "Returns the list of members that comprise the Global datastore.")
	})
	elasticacheCmd.AddCommand(elasticache_describeGlobalReplicationGroupsCmd)
}
