package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appsync_updateDataSourceCmd = &cobra.Command{
	Use:   "update-data-source",
	Short: "Updates a `DataSource` object.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appsync_updateDataSourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appsync_updateDataSourceCmd).Standalone()

		appsync_updateDataSourceCmd.Flags().String("api-id", "", "The API ID.")
		appsync_updateDataSourceCmd.Flags().String("description", "", "The new description for the data source.")
		appsync_updateDataSourceCmd.Flags().String("dynamodb-config", "", "The new Amazon DynamoDB configuration.")
		appsync_updateDataSourceCmd.Flags().String("elasticsearch-config", "", "The new OpenSearch configuration.")
		appsync_updateDataSourceCmd.Flags().String("event-bridge-config", "", "The new Amazon EventBridge settings.")
		appsync_updateDataSourceCmd.Flags().String("http-config", "", "The new HTTP endpoint configuration.")
		appsync_updateDataSourceCmd.Flags().String("lambda-config", "", "The new Lambda configuration.")
		appsync_updateDataSourceCmd.Flags().String("metrics-config", "", "Enables or disables enhanced data source metrics for specified data sources.")
		appsync_updateDataSourceCmd.Flags().String("name", "", "The new name for the data source.")
		appsync_updateDataSourceCmd.Flags().String("open-search-service-config", "", "The new OpenSearch configuration.")
		appsync_updateDataSourceCmd.Flags().String("relational-database-config", "", "The new relational database configuration.")
		appsync_updateDataSourceCmd.Flags().String("service-role-arn", "", "The new service role Amazon Resource Name (ARN) for the data source.")
		appsync_updateDataSourceCmd.Flags().String("type", "", "The new data source type.")
		appsync_updateDataSourceCmd.MarkFlagRequired("api-id")
		appsync_updateDataSourceCmd.MarkFlagRequired("name")
		appsync_updateDataSourceCmd.MarkFlagRequired("type")
	})
	appsyncCmd.AddCommand(appsync_updateDataSourceCmd)
}
