package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticache_testFailoverCmd = &cobra.Command{
	Use:   "test-failover",
	Short: "Represents the input of a `TestFailover` operation which tests automatic failover on a specified node group (called shard in the console) in a replication group (called cluster in the console).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticache_testFailoverCmd).Standalone()

	elasticache_testFailoverCmd.Flags().String("node-group-id", "", "The name of the node group (called shard in the console) in this replication group on which automatic failover is to be tested.")
	elasticache_testFailoverCmd.Flags().String("replication-group-id", "", "The name of the replication group (console: cluster) whose automatic failover is being tested by this operation.")
	elasticache_testFailoverCmd.MarkFlagRequired("node-group-id")
	elasticache_testFailoverCmd.MarkFlagRequired("replication-group-id")
	elasticacheCmd.AddCommand(elasticache_testFailoverCmd)
}
