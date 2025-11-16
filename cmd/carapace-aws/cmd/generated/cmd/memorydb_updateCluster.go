package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var memorydb_updateClusterCmd = &cobra.Command{
	Use:   "update-cluster",
	Short: "Modifies the settings for a cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(memorydb_updateClusterCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(memorydb_updateClusterCmd).Standalone()

		memorydb_updateClusterCmd.Flags().String("aclname", "", "The Access Control List that is associated with the cluster.")
		memorydb_updateClusterCmd.Flags().String("cluster-name", "", "The name of the cluster to update.")
		memorydb_updateClusterCmd.Flags().String("description", "", "The description of the cluster to update.")
		memorydb_updateClusterCmd.Flags().String("engine", "", "The name of the engine to be used for the cluster.")
		memorydb_updateClusterCmd.Flags().String("engine-version", "", "The upgraded version of the engine to be run on the nodes.")
		memorydb_updateClusterCmd.Flags().String("ip-discovery", "", "The mechanism for discovering IP addresses for the cluster discovery protocol.")
		memorydb_updateClusterCmd.Flags().String("maintenance-window", "", "Specifies the weekly time range during which maintenance on the cluster is performed.")
		memorydb_updateClusterCmd.Flags().String("node-type", "", "A valid node type that you want to scale this cluster up or down to.")
		memorydb_updateClusterCmd.Flags().String("parameter-group-name", "", "The name of the parameter group to update.")
		memorydb_updateClusterCmd.Flags().String("replica-configuration", "", "The number of replicas that will reside in each shard.")
		memorydb_updateClusterCmd.Flags().String("security-group-ids", "", "The SecurityGroupIds to update.")
		memorydb_updateClusterCmd.Flags().String("shard-configuration", "", "The number of shards in the cluster.")
		memorydb_updateClusterCmd.Flags().String("snapshot-retention-limit", "", "The number of days for which MemoryDB retains automatic cluster snapshots before deleting them.")
		memorydb_updateClusterCmd.Flags().String("snapshot-window", "", "The daily time range (in UTC) during which MemoryDB begins taking a daily snapshot of your cluster.")
		memorydb_updateClusterCmd.Flags().String("sns-topic-arn", "", "The SNS topic ARN to update.")
		memorydb_updateClusterCmd.Flags().String("sns-topic-status", "", "The status of the Amazon SNS notification topic.")
		memorydb_updateClusterCmd.MarkFlagRequired("cluster-name")
	})
	memorydbCmd.AddCommand(memorydb_updateClusterCmd)
}
