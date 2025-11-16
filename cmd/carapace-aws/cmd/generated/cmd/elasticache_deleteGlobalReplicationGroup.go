package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticache_deleteGlobalReplicationGroupCmd = &cobra.Command{
	Use:   "delete-global-replication-group",
	Short: "Deleting a Global datastore is a two-step process:\n\n- First, you must [DisassociateGlobalReplicationGroup]() to remove the secondary clusters in the Global datastore.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticache_deleteGlobalReplicationGroupCmd).Standalone()

	elasticache_deleteGlobalReplicationGroupCmd.Flags().String("global-replication-group-id", "", "The name of the Global datastore")
	elasticache_deleteGlobalReplicationGroupCmd.Flags().Bool("no-retain-primary-replication-group", false, "The primary replication group is retained as a standalone replication group.")
	elasticache_deleteGlobalReplicationGroupCmd.Flags().Bool("retain-primary-replication-group", false, "The primary replication group is retained as a standalone replication group.")
	elasticache_deleteGlobalReplicationGroupCmd.MarkFlagRequired("global-replication-group-id")
	elasticache_deleteGlobalReplicationGroupCmd.Flag("no-retain-primary-replication-group").Hidden = true
	elasticache_deleteGlobalReplicationGroupCmd.MarkFlagRequired("no-retain-primary-replication-group")
	elasticache_deleteGlobalReplicationGroupCmd.MarkFlagRequired("retain-primary-replication-group")
	elasticacheCmd.AddCommand(elasticache_deleteGlobalReplicationGroupCmd)
}
