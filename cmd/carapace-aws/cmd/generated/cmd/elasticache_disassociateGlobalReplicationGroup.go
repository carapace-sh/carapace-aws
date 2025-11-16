package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticache_disassociateGlobalReplicationGroupCmd = &cobra.Command{
	Use:   "disassociate-global-replication-group",
	Short: "Remove a secondary cluster from the Global datastore using the Global datastore name.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticache_disassociateGlobalReplicationGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elasticache_disassociateGlobalReplicationGroupCmd).Standalone()

		elasticache_disassociateGlobalReplicationGroupCmd.Flags().String("global-replication-group-id", "", "The name of the Global datastore")
		elasticache_disassociateGlobalReplicationGroupCmd.Flags().String("replication-group-id", "", "The name of the secondary cluster you wish to remove from the Global datastore")
		elasticache_disassociateGlobalReplicationGroupCmd.Flags().String("replication-group-region", "", "The Amazon region of secondary cluster you wish to remove from the Global datastore")
		elasticache_disassociateGlobalReplicationGroupCmd.MarkFlagRequired("global-replication-group-id")
		elasticache_disassociateGlobalReplicationGroupCmd.MarkFlagRequired("replication-group-id")
		elasticache_disassociateGlobalReplicationGroupCmd.MarkFlagRequired("replication-group-region")
	})
	elasticacheCmd.AddCommand(elasticache_disassociateGlobalReplicationGroupCmd)
}
