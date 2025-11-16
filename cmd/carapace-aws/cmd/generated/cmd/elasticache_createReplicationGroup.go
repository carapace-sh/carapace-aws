package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticache_createReplicationGroupCmd = &cobra.Command{
	Use:   "create-replication-group",
	Short: "Creates a Valkey or Redis OSS (cluster mode disabled) or a Valkey or Redis OSS (cluster mode enabled) replication group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticache_createReplicationGroupCmd).Standalone()

	elasticache_createReplicationGroupCmd.Flags().String("at-rest-encryption-enabled", "", "A flag that enables encryption at rest when set to `true`.")
	elasticache_createReplicationGroupCmd.Flags().String("auth-token", "", "**Reserved parameter.** The password used to access a password protected server.")
	elasticache_createReplicationGroupCmd.Flags().String("auto-minor-version-upgrade", "", "If you are running Valkey 7.2 and above or Redis OSS engine version 6.0 and above, set this parameter to yes to opt-in to the next auto minor version upgrade campaign.")
	elasticache_createReplicationGroupCmd.Flags().String("automatic-failover-enabled", "", "Specifies whether a read-only replica is automatically promoted to read/write primary if the existing primary fails.")
	elasticache_createReplicationGroupCmd.Flags().String("cache-node-type", "", "The compute and memory capacity of the nodes in the node group (shard).")
	elasticache_createReplicationGroupCmd.Flags().String("cache-parameter-group-name", "", "The name of the parameter group to associate with this replication group.")
	elasticache_createReplicationGroupCmd.Flags().String("cache-security-group-names", "", "A list of cache security group names to associate with this replication group.")
	elasticache_createReplicationGroupCmd.Flags().String("cache-subnet-group-name", "", "The name of the cache subnet group to be used for the replication group.")
	elasticache_createReplicationGroupCmd.Flags().String("cluster-mode", "", "Enabled or Disabled.")
	elasticache_createReplicationGroupCmd.Flags().String("data-tiering-enabled", "", "Enables data tiering.")
	elasticache_createReplicationGroupCmd.Flags().String("engine", "", "The name of the cache engine to be used for the clusters in this replication group.")
	elasticache_createReplicationGroupCmd.Flags().String("engine-version", "", "The version number of the cache engine to be used for the clusters in this replication group.")
	elasticache_createReplicationGroupCmd.Flags().String("global-replication-group-id", "", "The name of the Global datastore")
	elasticache_createReplicationGroupCmd.Flags().String("ip-discovery", "", "The network type you choose when creating a replication group, either `ipv4` | `ipv6`.")
	elasticache_createReplicationGroupCmd.Flags().String("kms-key-id", "", "The ID of the KMS key used to encrypt the disk in the cluster.")
	elasticache_createReplicationGroupCmd.Flags().String("log-delivery-configurations", "", "Specifies the destination, format and type of the logs.")
	elasticache_createReplicationGroupCmd.Flags().String("multi-azenabled", "", "A flag indicating if you have Multi-AZ enabled to enhance fault tolerance.")
	elasticache_createReplicationGroupCmd.Flags().String("network-type", "", "Must be either `ipv4` | `ipv6` | `dual_stack`.")
	elasticache_createReplicationGroupCmd.Flags().String("node-group-configuration", "", "A list of node group (shard) configuration options.")
	elasticache_createReplicationGroupCmd.Flags().String("notification-topic-arn", "", "The Amazon Resource Name (ARN) of the Amazon Simple Notification Service (SNS) topic to which notifications are sent.")
	elasticache_createReplicationGroupCmd.Flags().String("num-cache-clusters", "", "The number of clusters this replication group initially has.")
	elasticache_createReplicationGroupCmd.Flags().String("num-node-groups", "", "An optional parameter that specifies the number of node groups (shards) for this Valkey or Redis OSS (cluster mode enabled) replication group.")
	elasticache_createReplicationGroupCmd.Flags().String("port", "", "The port number on which each member of the replication group accepts connections.")
	elasticache_createReplicationGroupCmd.Flags().String("preferred-cache-cluster-azs", "", "A list of EC2 Availability Zones in which the replication group's clusters are created.")
	elasticache_createReplicationGroupCmd.Flags().String("preferred-maintenance-window", "", "Specifies the weekly time range during which maintenance on the cluster is performed.")
	elasticache_createReplicationGroupCmd.Flags().String("primary-cluster-id", "", "The identifier of the cluster that serves as the primary for this replication group.")
	elasticache_createReplicationGroupCmd.Flags().String("replicas-per-node-group", "", "An optional parameter that specifies the number of replica nodes in each node group (shard).")
	elasticache_createReplicationGroupCmd.Flags().String("replication-group-description", "", "A user-created description for the replication group.")
	elasticache_createReplicationGroupCmd.Flags().String("replication-group-id", "", "The replication group identifier.")
	elasticache_createReplicationGroupCmd.Flags().String("security-group-ids", "", "One or more Amazon VPC security groups associated with this replication group.")
	elasticache_createReplicationGroupCmd.Flags().String("serverless-cache-snapshot-name", "", "The name of the snapshot used to create a replication group.")
	elasticache_createReplicationGroupCmd.Flags().String("snapshot-arns", "", "A list of Amazon Resource Names (ARN) that uniquely identify the Valkey or Redis OSS RDB snapshot files stored in Amazon S3.")
	elasticache_createReplicationGroupCmd.Flags().String("snapshot-name", "", "The name of a snapshot from which to restore data into the new replication group.")
	elasticache_createReplicationGroupCmd.Flags().String("snapshot-retention-limit", "", "The number of days for which ElastiCache retains automatic snapshots before deleting them.")
	elasticache_createReplicationGroupCmd.Flags().String("snapshot-window", "", "The daily time range (in UTC) during which ElastiCache begins taking a daily snapshot of your node group (shard).")
	elasticache_createReplicationGroupCmd.Flags().String("tags", "", "A list of tags to be added to this resource.")
	elasticache_createReplicationGroupCmd.Flags().String("transit-encryption-enabled", "", "A flag that enables in-transit encryption when set to `true`.")
	elasticache_createReplicationGroupCmd.Flags().String("transit-encryption-mode", "", "A setting that allows you to migrate your clients to use in-transit encryption, with no downtime.")
	elasticache_createReplicationGroupCmd.Flags().String("user-group-ids", "", "The user group to associate with the replication group.")
	elasticache_createReplicationGroupCmd.MarkFlagRequired("replication-group-description")
	elasticache_createReplicationGroupCmd.MarkFlagRequired("replication-group-id")
	elasticacheCmd.AddCommand(elasticache_createReplicationGroupCmd)
}
