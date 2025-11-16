package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appsync_createDataSourceCmd = &cobra.Command{
	Use:   "create-data-source",
	Short: "Creates a `DataSource` object.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appsync_createDataSourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appsync_createDataSourceCmd).Standalone()

		appsync_createDataSourceCmd.Flags().String("api-id", "", "The API ID for the GraphQL API for the `DataSource`.")
		appsync_createDataSourceCmd.Flags().String("description", "", "A description of the `DataSource`.")
		appsync_createDataSourceCmd.Flags().String("dynamodb-config", "", "Amazon DynamoDB settings.")
		appsync_createDataSourceCmd.Flags().String("elasticsearch-config", "", "Amazon OpenSearch Service settings.")
		appsync_createDataSourceCmd.Flags().String("event-bridge-config", "", "Amazon EventBridge settings.")
		appsync_createDataSourceCmd.Flags().String("http-config", "", "HTTP endpoint settings.")
		appsync_createDataSourceCmd.Flags().String("lambda-config", "", "Lambda settings.")
		appsync_createDataSourceCmd.Flags().String("metrics-config", "", "Enables or disables enhanced data source metrics for specified data sources.")
		appsync_createDataSourceCmd.Flags().String("name", "", "A user-supplied name for the `DataSource`.")
		appsync_createDataSourceCmd.Flags().String("open-search-service-config", "", "Amazon OpenSearch Service settings.")
		appsync_createDataSourceCmd.Flags().String("relational-database-config", "", "Relational database settings.")
		appsync_createDataSourceCmd.Flags().String("service-role-arn", "", "The Identity and Access Management (IAM) service role Amazon Resource Name (ARN) for the data source.")
		appsync_createDataSourceCmd.Flags().String("type", "", "The type of the `DataSource`.")
		appsync_createDataSourceCmd.MarkFlagRequired("api-id")
		appsync_createDataSourceCmd.MarkFlagRequired("name")
		appsync_createDataSourceCmd.MarkFlagRequired("type")
	})
	appsyncCmd.AddCommand(appsync_createDataSourceCmd)
}
