package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appsync_updateResolverCmd = &cobra.Command{
	Use:   "update-resolver",
	Short: "Updates a `Resolver` object.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appsync_updateResolverCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appsync_updateResolverCmd).Standalone()

		appsync_updateResolverCmd.Flags().String("api-id", "", "The API ID.")
		appsync_updateResolverCmd.Flags().String("caching-config", "", "The caching configuration for the resolver.")
		appsync_updateResolverCmd.Flags().String("code", "", "The `resolver` code that contains the request and response functions.")
		appsync_updateResolverCmd.Flags().String("data-source-name", "", "The new data source name.")
		appsync_updateResolverCmd.Flags().String("field-name", "", "The new field name.")
		appsync_updateResolverCmd.Flags().String("kind", "", "The resolver type.")
		appsync_updateResolverCmd.Flags().String("max-batch-size", "", "The maximum batching size for a resolver.")
		appsync_updateResolverCmd.Flags().String("metrics-config", "", "Enables or disables enhanced resolver metrics for specified resolvers.")
		appsync_updateResolverCmd.Flags().String("pipeline-config", "", "The `PipelineConfig`.")
		appsync_updateResolverCmd.Flags().String("request-mapping-template", "", "The new request mapping template.")
		appsync_updateResolverCmd.Flags().String("response-mapping-template", "", "The new response mapping template.")
		appsync_updateResolverCmd.Flags().String("runtime", "", "")
		appsync_updateResolverCmd.Flags().String("sync-config", "", "The `SyncConfig` for a resolver attached to a versioned data source.")
		appsync_updateResolverCmd.Flags().String("type-name", "", "The new type name.")
		appsync_updateResolverCmd.MarkFlagRequired("api-id")
		appsync_updateResolverCmd.MarkFlagRequired("field-name")
		appsync_updateResolverCmd.MarkFlagRequired("type-name")
	})
	appsyncCmd.AddCommand(appsync_updateResolverCmd)
}
