package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticache_deleteServerlessCacheCmd = &cobra.Command{
	Use:   "delete-serverless-cache",
	Short: "Deletes a specified existing serverless cache.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticache_deleteServerlessCacheCmd).Standalone()

	elasticache_deleteServerlessCacheCmd.Flags().String("final-snapshot-name", "", "Name of the final snapshot to be taken before the serverless cache is deleted.")
	elasticache_deleteServerlessCacheCmd.Flags().String("serverless-cache-name", "", "The identifier of the serverless cache to be deleted.")
	elasticache_deleteServerlessCacheCmd.MarkFlagRequired("serverless-cache-name")
	elasticacheCmd.AddCommand(elasticache_deleteServerlessCacheCmd)
}
