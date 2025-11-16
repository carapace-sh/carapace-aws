package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appsync_createResolverCmd = &cobra.Command{
	Use:   "create-resolver",
	Short: "Creates a `Resolver` object.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appsync_createResolverCmd).Standalone()

	appsync_createResolverCmd.Flags().String("api-id", "", "The ID for the GraphQL API for which the resolver is being created.")
	appsync_createResolverCmd.Flags().String("caching-config", "", "The caching configuration for the resolver.")
	appsync_createResolverCmd.Flags().String("code", "", "The `resolver` code that contains the request and response functions.")
	appsync_createResolverCmd.Flags().String("data-source-name", "", "The name of the data source for which the resolver is being created.")
	appsync_createResolverCmd.Flags().String("field-name", "", "The name of the field to attach the resolver to.")
	appsync_createResolverCmd.Flags().String("kind", "", "The resolver type.")
	appsync_createResolverCmd.Flags().String("max-batch-size", "", "The maximum batching size for a resolver.")
	appsync_createResolverCmd.Flags().String("metrics-config", "", "Enables or disables enhanced resolver metrics for specified resolvers.")
	appsync_createResolverCmd.Flags().String("pipeline-config", "", "The `PipelineConfig`.")
	appsync_createResolverCmd.Flags().String("request-mapping-template", "", "The mapping template to use for requests.")
	appsync_createResolverCmd.Flags().String("response-mapping-template", "", "The mapping template to use for responses from the data source.")
	appsync_createResolverCmd.Flags().String("runtime", "", "")
	appsync_createResolverCmd.Flags().String("sync-config", "", "The `SyncConfig` for a resolver attached to a versioned data source.")
	appsync_createResolverCmd.Flags().String("type-name", "", "The name of the `Type`.")
	appsync_createResolverCmd.MarkFlagRequired("api-id")
	appsync_createResolverCmd.MarkFlagRequired("field-name")
	appsync_createResolverCmd.MarkFlagRequired("type-name")
	appsyncCmd.AddCommand(appsync_createResolverCmd)
}
