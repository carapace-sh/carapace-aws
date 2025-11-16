package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qbusiness_createDataSourceCmd = &cobra.Command{
	Use:   "create-data-source",
	Short: "Creates a data source connector for an Amazon Q Business application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qbusiness_createDataSourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qbusiness_createDataSourceCmd).Standalone()

		qbusiness_createDataSourceCmd.Flags().String("application-id", "", "The identifier of the Amazon Q Business application the data source will be attached to.")
		qbusiness_createDataSourceCmd.Flags().String("client-token", "", "A token you provide to identify a request to create a data source connector.")
		qbusiness_createDataSourceCmd.Flags().String("configuration", "", "Configuration information to connect your data source repository to Amazon Q Business.")
		qbusiness_createDataSourceCmd.Flags().String("description", "", "A description for the data source connector.")
		qbusiness_createDataSourceCmd.Flags().String("display-name", "", "A name for the data source connector.")
		qbusiness_createDataSourceCmd.Flags().String("document-enrichment-configuration", "", "")
		qbusiness_createDataSourceCmd.Flags().String("index-id", "", "The identifier of the index that you want to use with the data source connector.")
		qbusiness_createDataSourceCmd.Flags().String("media-extraction-configuration", "", "The configuration for extracting information from media in documents during ingestion.")
		qbusiness_createDataSourceCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of an IAM role with permission to access the data source and required resources.")
		qbusiness_createDataSourceCmd.Flags().String("sync-schedule", "", "Sets the frequency for Amazon Q Business to check the documents in your data source repository and update your index.")
		qbusiness_createDataSourceCmd.Flags().String("tags", "", "A list of key-value pairs that identify or categorize the data source connector.")
		qbusiness_createDataSourceCmd.Flags().String("vpc-configuration", "", "Configuration information for an Amazon VPC (Virtual Private Cloud) to connect to your data source.")
		qbusiness_createDataSourceCmd.MarkFlagRequired("application-id")
		qbusiness_createDataSourceCmd.MarkFlagRequired("configuration")
		qbusiness_createDataSourceCmd.MarkFlagRequired("display-name")
		qbusiness_createDataSourceCmd.MarkFlagRequired("index-id")
	})
	qbusinessCmd.AddCommand(qbusiness_createDataSourceCmd)
}
