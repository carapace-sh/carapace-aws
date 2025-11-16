package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticache_testMigrationCmd = &cobra.Command{
	Use:   "test-migration",
	Short: "Async API to test connection between source and target replication group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticache_testMigrationCmd).Standalone()

	elasticache_testMigrationCmd.Flags().String("customer-node-endpoint-list", "", "List of endpoints from which data should be migrated.")
	elasticache_testMigrationCmd.Flags().String("replication-group-id", "", "The ID of the replication group to which data is to be migrated.")
	elasticache_testMigrationCmd.MarkFlagRequired("customer-node-endpoint-list")
	elasticache_testMigrationCmd.MarkFlagRequired("replication-group-id")
	elasticacheCmd.AddCommand(elasticache_testMigrationCmd)
}
