package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var memorydb_createSnapshotCmd = &cobra.Command{
	Use:   "create-snapshot",
	Short: "Creates a copy of an entire cluster at a specific moment in time.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(memorydb_createSnapshotCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(memorydb_createSnapshotCmd).Standalone()

		memorydb_createSnapshotCmd.Flags().String("cluster-name", "", "The snapshot is created from this cluster.")
		memorydb_createSnapshotCmd.Flags().String("kms-key-id", "", "The ID of the KMS key used to encrypt the snapshot.")
		memorydb_createSnapshotCmd.Flags().String("snapshot-name", "", "A name for the snapshot being created.")
		memorydb_createSnapshotCmd.Flags().String("tags", "", "A list of tags to be added to this resource.")
		memorydb_createSnapshotCmd.MarkFlagRequired("cluster-name")
		memorydb_createSnapshotCmd.MarkFlagRequired("snapshot-name")
	})
	memorydbCmd.AddCommand(memorydb_createSnapshotCmd)
}
