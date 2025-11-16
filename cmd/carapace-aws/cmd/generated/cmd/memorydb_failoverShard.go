package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var memorydb_failoverShardCmd = &cobra.Command{
	Use:   "failover-shard",
	Short: "Used to failover a shard.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(memorydb_failoverShardCmd).Standalone()

	memorydb_failoverShardCmd.Flags().String("cluster-name", "", "The cluster being failed over.")
	memorydb_failoverShardCmd.Flags().String("shard-name", "", "The name of the shard.")
	memorydb_failoverShardCmd.MarkFlagRequired("cluster-name")
	memorydb_failoverShardCmd.MarkFlagRequired("shard-name")
	memorydbCmd.AddCommand(memorydb_failoverShardCmd)
}
