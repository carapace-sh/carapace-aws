package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticache_startMigrationCmd = &cobra.Command{
	Use:   "start-migration",
	Short: "Start the migration of data.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticache_startMigrationCmd).Standalone()

	elasticache_startMigrationCmd.Flags().String("customer-node-endpoint-list", "", "List of endpoints from which data should be migrated.")
	elasticache_startMigrationCmd.Flags().String("replication-group-id", "", "The ID of the replication group to which data should be migrated.")
	elasticache_startMigrationCmd.MarkFlagRequired("customer-node-endpoint-list")
	elasticache_startMigrationCmd.MarkFlagRequired("replication-group-id")
	elasticacheCmd.AddCommand(elasticache_startMigrationCmd)
}
