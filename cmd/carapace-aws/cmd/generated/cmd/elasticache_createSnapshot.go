package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticache_createSnapshotCmd = &cobra.Command{
	Use:   "create-snapshot",
	Short: "Creates a copy of an entire cluster or replication group at a specific moment in time.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticache_createSnapshotCmd).Standalone()

	elasticache_createSnapshotCmd.Flags().String("cache-cluster-id", "", "The identifier of an existing cluster.")
	elasticache_createSnapshotCmd.Flags().String("kms-key-id", "", "The ID of the KMS key used to encrypt the snapshot.")
	elasticache_createSnapshotCmd.Flags().String("replication-group-id", "", "The identifier of an existing replication group.")
	elasticache_createSnapshotCmd.Flags().String("snapshot-name", "", "A name for the snapshot being created.")
	elasticache_createSnapshotCmd.Flags().String("tags", "", "A list of tags to be added to this resource.")
	elasticache_createSnapshotCmd.MarkFlagRequired("snapshot-name")
	elasticacheCmd.AddCommand(elasticache_createSnapshotCmd)
}
