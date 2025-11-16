package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticache_completeMigrationCmd = &cobra.Command{
	Use:   "complete-migration",
	Short: "Complete the migration of data.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticache_completeMigrationCmd).Standalone()

	elasticache_completeMigrationCmd.Flags().Bool("force", false, "Forces the migration to stop without ensuring that data is in sync.")
	elasticache_completeMigrationCmd.Flags().Bool("no-force", false, "Forces the migration to stop without ensuring that data is in sync.")
	elasticache_completeMigrationCmd.Flags().String("replication-group-id", "", "The ID of the replication group to which data is being migrated.")
	elasticache_completeMigrationCmd.Flag("no-force").Hidden = true
	elasticache_completeMigrationCmd.MarkFlagRequired("replication-group-id")
	elasticacheCmd.AddCommand(elasticache_completeMigrationCmd)
}
