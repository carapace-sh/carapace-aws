package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticache_deleteServerlessCacheSnapshotCmd = &cobra.Command{
	Use:   "delete-serverless-cache-snapshot",
	Short: "Deletes an existing serverless cache snapshot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticache_deleteServerlessCacheSnapshotCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elasticache_deleteServerlessCacheSnapshotCmd).Standalone()

		elasticache_deleteServerlessCacheSnapshotCmd.Flags().String("serverless-cache-snapshot-name", "", "Idenfitier of the snapshot to be deleted.")
		elasticache_deleteServerlessCacheSnapshotCmd.MarkFlagRequired("serverless-cache-snapshot-name")
	})
	elasticacheCmd.AddCommand(elasticache_deleteServerlessCacheSnapshotCmd)
}
