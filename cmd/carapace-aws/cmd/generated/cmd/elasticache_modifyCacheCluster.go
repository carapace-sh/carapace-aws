package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticache_modifyCacheClusterCmd = &cobra.Command{
	Use:   "modify-cache-cluster",
	Short: "Modifies the settings for a cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticache_modifyCacheClusterCmd).Standalone()

	elasticache_modifyCacheClusterCmd.Flags().Bool("apply-immediately", false, "If `true`, this parameter causes the modifications in this request and any pending modifications to be applied, asynchronously and as soon as possible, regardless of the `PreferredMaintenanceWindow` setting for the cluster.")
	elasticache_modifyCacheClusterCmd.Flags().String("auth-token", "", "Reserved parameter.")
	elasticache_modifyCacheClusterCmd.Flags().String("auth-token-update-strategy", "", "Specifies the strategy to use to update the AUTH token.")
	elasticache_modifyCacheClusterCmd.Flags().String("auto-minor-version-upgrade", "", "If you are running Valkey 7.2 or Redis OSS engine version 6.0 or later, set this parameter to yes to opt-in to the next auto minor version upgrade campaign.")
	elasticache_modifyCacheClusterCmd.Flags().String("azmode", "", "Specifies whether the new nodes in this Memcached cluster are all created in a single Availability Zone or created across multiple Availability Zones.")
	elasticache_modifyCacheClusterCmd.Flags().String("cache-cluster-id", "", "The cluster identifier.")
	elasticache_modifyCacheClusterCmd.Flags().String("cache-node-ids-to-remove", "", "A list of cache node IDs to be removed.")
	elasticache_modifyCacheClusterCmd.Flags().String("cache-node-type", "", "A valid cache node type that you want to scale this cluster up to.")
	elasticache_modifyCacheClusterCmd.Flags().String("cache-parameter-group-name", "", "The name of the cache parameter group to apply to this cluster.")
	elasticache_modifyCacheClusterCmd.Flags().String("cache-security-group-names", "", "A list of cache security group names to authorize on this cluster.")
	elasticache_modifyCacheClusterCmd.Flags().String("engine", "", "The engine type used by the cache cluster.")
	elasticache_modifyCacheClusterCmd.Flags().String("engine-version", "", "The upgraded version of the cache engine to be run on the cache nodes.")
	elasticache_modifyCacheClusterCmd.Flags().String("ip-discovery", "", "The network type you choose when modifying a cluster, either `ipv4` | `ipv6`.")
	elasticache_modifyCacheClusterCmd.Flags().String("log-delivery-configurations", "", "Specifies the destination, format and type of the logs.")
	elasticache_modifyCacheClusterCmd.Flags().String("new-availability-zones", "", "This option is only supported on Memcached clusters.")
	elasticache_modifyCacheClusterCmd.Flags().Bool("no-apply-immediately", false, "If `true`, this parameter causes the modifications in this request and any pending modifications to be applied, asynchronously and as soon as possible, regardless of the `PreferredMaintenanceWindow` setting for the cluster.")
	elasticache_modifyCacheClusterCmd.Flags().String("notification-topic-arn", "", "The Amazon Resource Name (ARN) of the Amazon SNS topic to which notifications are sent.")
	elasticache_modifyCacheClusterCmd.Flags().String("notification-topic-status", "", "The status of the Amazon SNS notification topic.")
	elasticache_modifyCacheClusterCmd.Flags().String("num-cache-nodes", "", "The number of cache nodes that the cluster should have.")
	elasticache_modifyCacheClusterCmd.Flags().String("preferred-maintenance-window", "", "Specifies the weekly time range during which maintenance on the cluster is performed.")
	elasticache_modifyCacheClusterCmd.Flags().String("scale-config", "", "Configures horizontal or vertical scaling for Memcached clusters, specifying the scaling percentage and interval.")
	elasticache_modifyCacheClusterCmd.Flags().String("security-group-ids", "", "Specifies the VPC Security Groups associated with the cluster.")
	elasticache_modifyCacheClusterCmd.Flags().String("snapshot-retention-limit", "", "The number of days for which ElastiCache retains automatic cluster snapshots before deleting them.")
	elasticache_modifyCacheClusterCmd.Flags().String("snapshot-window", "", "The daily time range (in UTC) during which ElastiCache begins taking a daily snapshot of your cluster.")
	elasticache_modifyCacheClusterCmd.MarkFlagRequired("cache-cluster-id")
	elasticache_modifyCacheClusterCmd.Flag("no-apply-immediately").Hidden = true
	elasticacheCmd.AddCommand(elasticache_modifyCacheClusterCmd)
}
