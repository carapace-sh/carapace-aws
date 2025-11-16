package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appsync_updateApiCacheCmd = &cobra.Command{
	Use:   "update-api-cache",
	Short: "Updates the cache for the GraphQL API.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appsync_updateApiCacheCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appsync_updateApiCacheCmd).Standalone()

		appsync_updateApiCacheCmd.Flags().String("api-caching-behavior", "", "Caching behavior.")
		appsync_updateApiCacheCmd.Flags().String("api-id", "", "The GraphQL API ID.")
		appsync_updateApiCacheCmd.Flags().String("health-metrics-config", "", "Controls how cache health metrics will be emitted to CloudWatch.")
		appsync_updateApiCacheCmd.Flags().String("ttl", "", "TTL in seconds for cache entries.")
		appsync_updateApiCacheCmd.Flags().String("type", "", "The cache instance type.")
		appsync_updateApiCacheCmd.MarkFlagRequired("api-caching-behavior")
		appsync_updateApiCacheCmd.MarkFlagRequired("api-id")
		appsync_updateApiCacheCmd.MarkFlagRequired("ttl")
		appsync_updateApiCacheCmd.MarkFlagRequired("type")
	})
	appsyncCmd.AddCommand(appsync_updateApiCacheCmd)
}
