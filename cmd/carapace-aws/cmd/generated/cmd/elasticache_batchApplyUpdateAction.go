package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticache_batchApplyUpdateActionCmd = &cobra.Command{
	Use:   "batch-apply-update-action",
	Short: "Apply the service update.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticache_batchApplyUpdateActionCmd).Standalone()

	elasticache_batchApplyUpdateActionCmd.Flags().String("cache-cluster-ids", "", "The cache cluster IDs")
	elasticache_batchApplyUpdateActionCmd.Flags().String("replication-group-ids", "", "The replication group IDs")
	elasticache_batchApplyUpdateActionCmd.Flags().String("service-update-name", "", "The unique ID of the service update")
	elasticache_batchApplyUpdateActionCmd.MarkFlagRequired("service-update-name")
	elasticacheCmd.AddCommand(elasticache_batchApplyUpdateActionCmd)
}
