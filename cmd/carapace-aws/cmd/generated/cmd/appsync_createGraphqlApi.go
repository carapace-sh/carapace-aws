package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appsync_createGraphqlApiCmd = &cobra.Command{
	Use:   "create-graphql-api",
	Short: "Creates a `GraphqlApi` object.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appsync_createGraphqlApiCmd).Standalone()

	appsync_createGraphqlApiCmd.Flags().String("additional-authentication-providers", "", "A list of additional authentication providers for the `GraphqlApi` API.")
	appsync_createGraphqlApiCmd.Flags().String("api-type", "", "The value that indicates whether the GraphQL API is a standard API (`GRAPHQL`) or merged API (`MERGED`).")
	appsync_createGraphqlApiCmd.Flags().String("authentication-type", "", "The authentication type: API key, Identity and Access Management (IAM), OpenID Connect (OIDC), Amazon Cognito user pools, or Lambda.")
	appsync_createGraphqlApiCmd.Flags().String("enhanced-metrics-config", "", "The `enhancedMetricsConfig` object.")
	appsync_createGraphqlApiCmd.Flags().String("introspection-config", "", "Sets the value of the GraphQL API to enable (`ENABLED`) or disable (`DISABLED`) introspection.")
	appsync_createGraphqlApiCmd.Flags().String("lambda-authorizer-config", "", "Configuration for Lambda function authorization.")
	appsync_createGraphqlApiCmd.Flags().String("log-config", "", "The Amazon CloudWatch Logs configuration.")
	appsync_createGraphqlApiCmd.Flags().String("merged-api-execution-role-arn", "", "The Identity and Access Management service role ARN for a merged API.")
	appsync_createGraphqlApiCmd.Flags().String("name", "", "A user-supplied name for the `GraphqlApi`.")
	appsync_createGraphqlApiCmd.Flags().Bool("no-xray-enabled", false, "A flag indicating whether to use X-Ray tracing for the `GraphqlApi`.")
	appsync_createGraphqlApiCmd.Flags().String("open-idconnect-config", "", "The OIDC configuration.")
	appsync_createGraphqlApiCmd.Flags().String("owner-contact", "", "The owner contact information for an API resource.")
	appsync_createGraphqlApiCmd.Flags().String("query-depth-limit", "", "The maximum depth a query can have in a single request.")
	appsync_createGraphqlApiCmd.Flags().String("resolver-count-limit", "", "The maximum number of resolvers that can be invoked in a single request.")
	appsync_createGraphqlApiCmd.Flags().String("tags", "", "A `TagMap` object.")
	appsync_createGraphqlApiCmd.Flags().String("user-pool-config", "", "The Amazon Cognito user pool configuration.")
	appsync_createGraphqlApiCmd.Flags().String("visibility", "", "Sets the value of the GraphQL API to public (`GLOBAL`) or private (`PRIVATE`).")
	appsync_createGraphqlApiCmd.Flags().Bool("xray-enabled", false, "A flag indicating whether to use X-Ray tracing for the `GraphqlApi`.")
	appsync_createGraphqlApiCmd.MarkFlagRequired("authentication-type")
	appsync_createGraphqlApiCmd.MarkFlagRequired("name")
	appsync_createGraphqlApiCmd.Flag("no-xray-enabled").Hidden = true
	appsyncCmd.AddCommand(appsync_createGraphqlApiCmd)
}
