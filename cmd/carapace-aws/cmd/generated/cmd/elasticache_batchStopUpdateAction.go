package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticache_batchStopUpdateActionCmd = &cobra.Command{
	Use:   "batch-stop-update-action",
	Short: "Stop the service update.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticache_batchStopUpdateActionCmd).Standalone()

	elasticache_batchStopUpdateActionCmd.Flags().String("cache-cluster-ids", "", "The cache cluster IDs")
	elasticache_batchStopUpdateActionCmd.Flags().String("replication-group-ids", "", "The replication group IDs")
	elasticache_batchStopUpdateActionCmd.Flags().String("service-update-name", "", "The unique ID of the service update")
	elasticache_batchStopUpdateActionCmd.MarkFlagRequired("service-update-name")
	elasticacheCmd.AddCommand(elasticache_batchStopUpdateActionCmd)
}
