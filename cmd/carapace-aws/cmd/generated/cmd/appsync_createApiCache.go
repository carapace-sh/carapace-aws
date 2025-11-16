package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appsync_createApiCacheCmd = &cobra.Command{
	Use:   "create-api-cache",
	Short: "Creates a cache for the GraphQL API.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appsync_createApiCacheCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appsync_createApiCacheCmd).Standalone()

		appsync_createApiCacheCmd.Flags().String("api-caching-behavior", "", "Caching behavior.")
		appsync_createApiCacheCmd.Flags().String("api-id", "", "The GraphQL API ID.")
		appsync_createApiCacheCmd.Flags().Bool("at-rest-encryption-enabled", false, "At-rest encryption flag for cache.")
		appsync_createApiCacheCmd.Flags().String("health-metrics-config", "", "Controls how cache health metrics will be emitted to CloudWatch.")
		appsync_createApiCacheCmd.Flags().Bool("no-at-rest-encryption-enabled", false, "At-rest encryption flag for cache.")
		appsync_createApiCacheCmd.Flags().Bool("no-transit-encryption-enabled", false, "Transit encryption flag when connecting to cache.")
		appsync_createApiCacheCmd.Flags().Bool("transit-encryption-enabled", false, "Transit encryption flag when connecting to cache.")
		appsync_createApiCacheCmd.Flags().String("ttl", "", "TTL in seconds for cache entries.")
		appsync_createApiCacheCmd.Flags().String("type", "", "The cache instance type.")
		appsync_createApiCacheCmd.MarkFlagRequired("api-caching-behavior")
		appsync_createApiCacheCmd.MarkFlagRequired("api-id")
		appsync_createApiCacheCmd.Flag("no-at-rest-encryption-enabled").Hidden = true
		appsync_createApiCacheCmd.Flag("no-transit-encryption-enabled").Hidden = true
		appsync_createApiCacheCmd.MarkFlagRequired("ttl")
		appsync_createApiCacheCmd.MarkFlagRequired("type")
	})
	appsyncCmd.AddCommand(appsync_createApiCacheCmd)
}
