package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appsync_updateGraphqlApiCmd = &cobra.Command{
	Use:   "update-graphql-api",
	Short: "Updates a `GraphqlApi` object.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appsync_updateGraphqlApiCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appsync_updateGraphqlApiCmd).Standalone()

		appsync_updateGraphqlApiCmd.Flags().String("additional-authentication-providers", "", "A list of additional authentication providers for the `GraphqlApi` API.")
		appsync_updateGraphqlApiCmd.Flags().String("api-id", "", "The API ID.")
		appsync_updateGraphqlApiCmd.Flags().String("authentication-type", "", "The new authentication type for the `GraphqlApi` object.")
		appsync_updateGraphqlApiCmd.Flags().String("enhanced-metrics-config", "", "The `enhancedMetricsConfig` object.")
		appsync_updateGraphqlApiCmd.Flags().String("introspection-config", "", "Sets the value of the GraphQL API to enable (`ENABLED`) or disable (`DISABLED`) introspection.")
		appsync_updateGraphqlApiCmd.Flags().String("lambda-authorizer-config", "", "Configuration for Lambda function authorization.")
		appsync_updateGraphqlApiCmd.Flags().String("log-config", "", "The Amazon CloudWatch Logs configuration for the `GraphqlApi` object.")
		appsync_updateGraphqlApiCmd.Flags().String("merged-api-execution-role-arn", "", "The Identity and Access Management service role ARN for a merged API.")
		appsync_updateGraphqlApiCmd.Flags().String("name", "", "The new name for the `GraphqlApi` object.")
		appsync_updateGraphqlApiCmd.Flags().Bool("no-xray-enabled", false, "A flag indicating whether to use X-Ray tracing for the `GraphqlApi`.")
		appsync_updateGraphqlApiCmd.Flags().String("open-idconnect-config", "", "The OpenID Connect configuration for the `GraphqlApi` object.")
		appsync_updateGraphqlApiCmd.Flags().String("owner-contact", "", "The owner contact information for an API resource.")
		appsync_updateGraphqlApiCmd.Flags().String("query-depth-limit", "", "The maximum depth a query can have in a single request.")
		appsync_updateGraphqlApiCmd.Flags().String("resolver-count-limit", "", "The maximum number of resolvers that can be invoked in a single request.")
		appsync_updateGraphqlApiCmd.Flags().String("user-pool-config", "", "The new Amazon Cognito user pool configuration for the `~GraphqlApi` object.")
		appsync_updateGraphqlApiCmd.Flags().Bool("xray-enabled", false, "A flag indicating whether to use X-Ray tracing for the `GraphqlApi`.")
		appsync_updateGraphqlApiCmd.MarkFlagRequired("api-id")
		appsync_updateGraphqlApiCmd.MarkFlagRequired("authentication-type")
		appsync_updateGraphqlApiCmd.MarkFlagRequired("name")
		appsync_updateGraphqlApiCmd.Flag("no-xray-enabled").Hidden = true
	})
	appsyncCmd.AddCommand(appsync_updateGraphqlApiCmd)
}
