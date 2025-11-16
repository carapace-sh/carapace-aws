package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticache_createCacheClusterCmd = &cobra.Command{
	Use:   "create-cache-cluster",
	Short: "Creates a cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticache_createCacheClusterCmd).Standalone()

	elasticache_createCacheClusterCmd.Flags().String("auth-token", "", "**Reserved parameter.** The password used to access a password protected server.")
	elasticache_createCacheClusterCmd.Flags().String("auto-minor-version-upgrade", "", "If you are running Valkey 7.2 and above or Redis OSS engine version 6.0 and above, set this parameter to yes to opt-in to the next auto minor version upgrade campaign.")
	elasticache_createCacheClusterCmd.Flags().String("azmode", "", "Specifies whether the nodes in this Memcached cluster are created in a single Availability Zone or created across multiple Availability Zones in the cluster's region.")
	elasticache_createCacheClusterCmd.Flags().String("cache-cluster-id", "", "The node group (shard) identifier.")
	elasticache_createCacheClusterCmd.Flags().String("cache-node-type", "", "The compute and memory capacity of the nodes in the node group (shard).")
	elasticache_createCacheClusterCmd.Flags().String("cache-parameter-group-name", "", "The name of the parameter group to associate with this cluster.")
	elasticache_createCacheClusterCmd.Flags().String("cache-security-group-names", "", "A list of security group names to associate with this cluster.")
	elasticache_createCacheClusterCmd.Flags().String("cache-subnet-group-name", "", "The name of the subnet group to be used for the cluster.")
	elasticache_createCacheClusterCmd.Flags().String("engine", "", "The name of the cache engine to be used for this cluster.")
	elasticache_createCacheClusterCmd.Flags().String("engine-version", "", "The version number of the cache engine to be used for this cluster.")
	elasticache_createCacheClusterCmd.Flags().String("ip-discovery", "", "The network type you choose when modifying a cluster, either `ipv4` | `ipv6`.")
	elasticache_createCacheClusterCmd.Flags().String("log-delivery-configurations", "", "Specifies the destination, format and type of the logs.")
	elasticache_createCacheClusterCmd.Flags().String("network-type", "", "Must be either `ipv4` | `ipv6` | `dual_stack`.")
	elasticache_createCacheClusterCmd.Flags().String("notification-topic-arn", "", "The Amazon Resource Name (ARN) of the Amazon Simple Notification Service (SNS) topic to which notifications are sent.")
	elasticache_createCacheClusterCmd.Flags().String("num-cache-nodes", "", "The initial number of cache nodes that the cluster has.")
	elasticache_createCacheClusterCmd.Flags().String("outpost-mode", "", "Specifies whether the nodes in the cluster are created in a single outpost or across multiple outposts.")
	elasticache_createCacheClusterCmd.Flags().String("port", "", "The port number on which each of the cache nodes accepts connections.")
	elasticache_createCacheClusterCmd.Flags().String("preferred-availability-zone", "", "The EC2 Availability Zone in which the cluster is created.")
	elasticache_createCacheClusterCmd.Flags().String("preferred-availability-zones", "", "A list of the Availability Zones in which cache nodes are created.")
	elasticache_createCacheClusterCmd.Flags().String("preferred-maintenance-window", "", "Specifies the weekly time range during which maintenance on the cluster is performed.")
	elasticache_createCacheClusterCmd.Flags().String("preferred-outpost-arn", "", "The outpost ARN in which the cache cluster is created.")
	elasticache_createCacheClusterCmd.Flags().String("preferred-outpost-arns", "", "The outpost ARNs in which the cache cluster is created.")
	elasticache_createCacheClusterCmd.Flags().String("replication-group-id", "", "The ID of the replication group to which this cluster should belong.")
	elasticache_createCacheClusterCmd.Flags().String("security-group-ids", "", "One or more VPC security groups associated with the cluster.")
	elasticache_createCacheClusterCmd.Flags().String("snapshot-arns", "", "A single-element string list containing an Amazon Resource Name (ARN) that uniquely identifies a Valkey or Redis OSS RDB snapshot file stored in Amazon S3.")
	elasticache_createCacheClusterCmd.Flags().String("snapshot-name", "", "The name of a Valkey or Redis OSS snapshot from which to restore data into the new node group (shard).")
	elasticache_createCacheClusterCmd.Flags().String("snapshot-retention-limit", "", "The number of days for which ElastiCache retains automatic snapshots before deleting them.")
	elasticache_createCacheClusterCmd.Flags().String("snapshot-window", "", "The daily time range (in UTC) during which ElastiCache begins taking a daily snapshot of your node group (shard).")
	elasticache_createCacheClusterCmd.Flags().String("tags", "", "A list of tags to be added to this resource.")
	elasticache_createCacheClusterCmd.Flags().String("transit-encryption-enabled", "", "A flag that enables in-transit encryption when set to true.")
	elasticache_createCacheClusterCmd.MarkFlagRequired("cache-cluster-id")
	elasticacheCmd.AddCommand(elasticache_createCacheClusterCmd)
}
