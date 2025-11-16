package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticache_deleteSnapshotCmd = &cobra.Command{
	Use:   "delete-snapshot",
	Short: "Deletes an existing snapshot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticache_deleteSnapshotCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elasticache_deleteSnapshotCmd).Standalone()

		elasticache_deleteSnapshotCmd.Flags().String("snapshot-name", "", "The name of the snapshot to be deleted.")
		elasticache_deleteSnapshotCmd.MarkFlagRequired("snapshot-name")
	})
	elasticacheCmd.AddCommand(elasticache_deleteSnapshotCmd)
}
