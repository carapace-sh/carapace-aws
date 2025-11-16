package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticache_modifyReplicationGroupCmd = &cobra.Command{
	Use:   "modify-replication-group",
	Short: "Modifies the settings for a replication group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticache_modifyReplicationGroupCmd).Standalone()

	elasticache_modifyReplicationGroupCmd.Flags().Bool("apply-immediately", false, "If `true`, this parameter causes the modifications in this request and any pending modifications to be applied, asynchronously and as soon as possible, regardless of the `PreferredMaintenanceWindow` setting for the replication group.")
	elasticache_modifyReplicationGroupCmd.Flags().String("auth-token", "", "Reserved parameter.")
	elasticache_modifyReplicationGroupCmd.Flags().String("auth-token-update-strategy", "", "Specifies the strategy to use to update the AUTH token.")
	elasticache_modifyReplicationGroupCmd.Flags().String("auto-minor-version-upgrade", "", "If you are running Valkey or Redis OSS engine version 6.0 or later, set this parameter to yes if you want to opt-in to the next auto minor version upgrade campaign.")
	elasticache_modifyReplicationGroupCmd.Flags().String("automatic-failover-enabled", "", "Determines whether a read replica is automatically promoted to read/write primary if the existing primary encounters a failure.")
	elasticache_modifyReplicationGroupCmd.Flags().String("cache-node-type", "", "A valid cache node type that you want to scale this replication group to.")
	elasticache_modifyReplicationGroupCmd.Flags().String("cache-parameter-group-name", "", "The name of the cache parameter group to apply to all of the clusters in this replication group.")
	elasticache_modifyReplicationGroupCmd.Flags().String("cache-security-group-names", "", "A list of cache security group names to authorize for the clusters in this replication group.")
	elasticache_modifyReplicationGroupCmd.Flags().String("cluster-mode", "", "Enabled or Disabled.")
	elasticache_modifyReplicationGroupCmd.Flags().String("engine", "", "Modifies the engine listed in a replication group message.")
	elasticache_modifyReplicationGroupCmd.Flags().String("engine-version", "", "The upgraded version of the cache engine to be run on the clusters in the replication group.")
	elasticache_modifyReplicationGroupCmd.Flags().String("ip-discovery", "", "The network type you choose when modifying a cluster, either `ipv4` | `ipv6`.")
	elasticache_modifyReplicationGroupCmd.Flags().String("log-delivery-configurations", "", "Specifies the destination, format and type of the logs.")
	elasticache_modifyReplicationGroupCmd.Flags().String("multi-azenabled", "", "A flag to indicate MultiAZ is enabled.")
	elasticache_modifyReplicationGroupCmd.Flags().Bool("no-apply-immediately", false, "If `true`, this parameter causes the modifications in this request and any pending modifications to be applied, asynchronously and as soon as possible, regardless of the `PreferredMaintenanceWindow` setting for the replication group.")
	elasticache_modifyReplicationGroupCmd.Flags().String("node-group-id", "", "Deprecated.")
	elasticache_modifyReplicationGroupCmd.Flags().String("notification-topic-arn", "", "The Amazon Resource Name (ARN) of the Amazon SNS topic to which notifications are sent.")
	elasticache_modifyReplicationGroupCmd.Flags().String("notification-topic-status", "", "The status of the Amazon SNS notification topic for the replication group.")
	elasticache_modifyReplicationGroupCmd.Flags().String("preferred-maintenance-window", "", "Specifies the weekly time range during which maintenance on the cluster is performed.")
	elasticache_modifyReplicationGroupCmd.Flags().String("primary-cluster-id", "", "For replication groups with a single primary, if this parameter is specified, ElastiCache promotes the specified cluster in the specified replication group to the primary role.")
	elasticache_modifyReplicationGroupCmd.Flags().String("remove-user-groups", "", "Removes the user group associated with this replication group.")
	elasticache_modifyReplicationGroupCmd.Flags().String("replication-group-description", "", "A description for the replication group.")
	elasticache_modifyReplicationGroupCmd.Flags().String("replication-group-id", "", "The identifier of the replication group to modify.")
	elasticache_modifyReplicationGroupCmd.Flags().String("security-group-ids", "", "Specifies the VPC Security Groups associated with the clusters in the replication group.")
	elasticache_modifyReplicationGroupCmd.Flags().String("snapshot-retention-limit", "", "The number of days for which ElastiCache retains automatic node group (shard) snapshots before deleting them.")
	elasticache_modifyReplicationGroupCmd.Flags().String("snapshot-window", "", "The daily time range (in UTC) during which ElastiCache begins taking a daily snapshot of the node group (shard) specified by `SnapshottingClusterId`.")
	elasticache_modifyReplicationGroupCmd.Flags().String("snapshotting-cluster-id", "", "The cluster ID that is used as the daily snapshot source for the replication group.")
	elasticache_modifyReplicationGroupCmd.Flags().String("transit-encryption-enabled", "", "A flag that enables in-transit encryption when set to true.")
	elasticache_modifyReplicationGroupCmd.Flags().String("transit-encryption-mode", "", "A setting that allows you to migrate your clients to use in-transit encryption, with no downtime.")
	elasticache_modifyReplicationGroupCmd.Flags().String("user-group-ids-to-add", "", "The ID of the user group you are associating with the replication group.")
	elasticache_modifyReplicationGroupCmd.Flags().String("user-group-ids-to-remove", "", "The ID of the user group to disassociate from the replication group, meaning the users in the group no longer can access the replication group.")
	elasticache_modifyReplicationGroupCmd.Flag("no-apply-immediately").Hidden = true
	elasticache_modifyReplicationGroupCmd.MarkFlagRequired("replication-group-id")
	elasticacheCmd.AddCommand(elasticache_modifyReplicationGroupCmd)
}
