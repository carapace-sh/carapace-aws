package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticache_describeSnapshotsCmd = &cobra.Command{
	Use:   "describe-snapshots",
	Short: "Returns information about cluster or replication group snapshots.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticache_describeSnapshotsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elasticache_describeSnapshotsCmd).Standalone()

		elasticache_describeSnapshotsCmd.Flags().String("cache-cluster-id", "", "A user-supplied cluster identifier.")
		elasticache_describeSnapshotsCmd.Flags().String("marker", "", "An optional marker returned from a prior request.")
		elasticache_describeSnapshotsCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
		elasticache_describeSnapshotsCmd.Flags().String("replication-group-id", "", "A user-supplied replication group identifier.")
		elasticache_describeSnapshotsCmd.Flags().String("show-node-group-config", "", "A Boolean value which if true, the node group (shard) configuration is included in the snapshot description.")
		elasticache_describeSnapshotsCmd.Flags().String("snapshot-name", "", "A user-supplied name of the snapshot.")
		elasticache_describeSnapshotsCmd.Flags().String("snapshot-source", "", "If set to `system`, the output shows snapshots that were automatically created by ElastiCache.")
	})
	elasticacheCmd.AddCommand(elasticache_describeSnapshotsCmd)
}
