package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var memorydb_createClusterCmd = &cobra.Command{
	Use:   "create-cluster",
	Short: "Creates a cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(memorydb_createClusterCmd).Standalone()

	memorydb_createClusterCmd.Flags().String("aclname", "", "The name of the Access Control List to associate with the cluster.")
	memorydb_createClusterCmd.Flags().String("auto-minor-version-upgrade", "", "When set to true, the cluster will automatically receive minor engine version upgrades after launch.")
	memorydb_createClusterCmd.Flags().String("cluster-name", "", "The name of the cluster.")
	memorydb_createClusterCmd.Flags().String("data-tiering", "", "Enables data tiering.")
	memorydb_createClusterCmd.Flags().String("description", "", "An optional description of the cluster.")
	memorydb_createClusterCmd.Flags().String("engine", "", "The name of the engine to be used for the cluster.")
	memorydb_createClusterCmd.Flags().String("engine-version", "", "The version number of the Redis OSS engine to be used for the cluster.")
	memorydb_createClusterCmd.Flags().String("ip-discovery", "", "The mechanism for discovering IP addresses for the cluster discovery protocol.")
	memorydb_createClusterCmd.Flags().String("kms-key-id", "", "The ID of the KMS key used to encrypt the cluster.")
	memorydb_createClusterCmd.Flags().String("maintenance-window", "", "Specifies the weekly time range during which maintenance on the cluster is performed.")
	memorydb_createClusterCmd.Flags().String("multi-region-cluster-name", "", "The name of the multi-Region cluster to be created.")
	memorydb_createClusterCmd.Flags().String("network-type", "", "Specifies the IP address type for the cluster.")
	memorydb_createClusterCmd.Flags().String("node-type", "", "The compute and memory capacity of the nodes in the cluster.")
	memorydb_createClusterCmd.Flags().String("num-replicas-per-shard", "", "The number of replicas to apply to each shard.")
	memorydb_createClusterCmd.Flags().String("num-shards", "", "The number of shards the cluster will contain.")
	memorydb_createClusterCmd.Flags().String("parameter-group-name", "", "The name of the parameter group associated with the cluster.")
	memorydb_createClusterCmd.Flags().String("port", "", "The port number on which each of the nodes accepts connections.")
	memorydb_createClusterCmd.Flags().String("security-group-ids", "", "A list of security group names to associate with this cluster.")
	memorydb_createClusterCmd.Flags().String("snapshot-arns", "", "A list of Amazon Resource Names (ARN) that uniquely identify the RDB snapshot files stored in Amazon S3.")
	memorydb_createClusterCmd.Flags().String("snapshot-name", "", "The name of a snapshot from which to restore data into the new cluster.")
	memorydb_createClusterCmd.Flags().String("snapshot-retention-limit", "", "The number of days for which MemoryDB retains automatic snapshots before deleting them.")
	memorydb_createClusterCmd.Flags().String("snapshot-window", "", "The daily time range (in UTC) during which MemoryDB begins taking a daily snapshot of your shard.")
	memorydb_createClusterCmd.Flags().String("sns-topic-arn", "", "The Amazon Resource Name (ARN) of the Amazon Simple Notification Service (SNS) topic to which notifications are sent.")
	memorydb_createClusterCmd.Flags().String("subnet-group-name", "", "The name of the subnet group to be used for the cluster.")
	memorydb_createClusterCmd.Flags().String("tags", "", "A list of tags to be added to this resource.")
	memorydb_createClusterCmd.Flags().String("tlsenabled", "", "A flag to enable in-transit encryption on the cluster.")
	memorydb_createClusterCmd.MarkFlagRequired("aclname")
	memorydb_createClusterCmd.MarkFlagRequired("cluster-name")
	memorydb_createClusterCmd.MarkFlagRequired("node-type")
	memorydbCmd.AddCommand(memorydb_createClusterCmd)
}
