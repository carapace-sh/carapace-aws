package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticache_describeUpdateActionsCmd = &cobra.Command{
	Use:   "describe-update-actions",
	Short: "Returns details of the update actions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticache_describeUpdateActionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elasticache_describeUpdateActionsCmd).Standalone()

		elasticache_describeUpdateActionsCmd.Flags().String("cache-cluster-ids", "", "The cache cluster IDs")
		elasticache_describeUpdateActionsCmd.Flags().String("engine", "", "The Elasticache engine to which the update applies.")
		elasticache_describeUpdateActionsCmd.Flags().String("marker", "", "An optional marker returned from a prior request.")
		elasticache_describeUpdateActionsCmd.Flags().String("max-records", "", "The maximum number of records to include in the response")
		elasticache_describeUpdateActionsCmd.Flags().String("replication-group-ids", "", "The replication group IDs")
		elasticache_describeUpdateActionsCmd.Flags().String("service-update-name", "", "The unique ID of the service update")
		elasticache_describeUpdateActionsCmd.Flags().String("service-update-status", "", "The status of the service update")
		elasticache_describeUpdateActionsCmd.Flags().String("service-update-time-range", "", "The range of time specified to search for service updates that are in available status")
		elasticache_describeUpdateActionsCmd.Flags().String("show-node-level-update-status", "", "Dictates whether to include node level update status in the response")
		elasticache_describeUpdateActionsCmd.Flags().String("update-action-status", "", "The status of the update action.")
	})
	elasticacheCmd.AddCommand(elasticache_describeUpdateActionsCmd)
}
