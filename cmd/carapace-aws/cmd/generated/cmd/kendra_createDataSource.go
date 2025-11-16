package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kendra_createDataSourceCmd = &cobra.Command{
	Use:   "create-data-source",
	Short: "Creates a data source connector that you want to use with an Amazon Kendra index.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kendra_createDataSourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kendra_createDataSourceCmd).Standalone()

		kendra_createDataSourceCmd.Flags().String("client-token", "", "A token that you provide to identify the request to create a data source connector.")
		kendra_createDataSourceCmd.Flags().String("configuration", "", "Configuration information to connect to your data source repository.")
		kendra_createDataSourceCmd.Flags().String("custom-document-enrichment-configuration", "", "Configuration information for altering document metadata and content during the document ingestion process.")
		kendra_createDataSourceCmd.Flags().String("description", "", "A description for the data source connector.")
		kendra_createDataSourceCmd.Flags().String("index-id", "", "The identifier of the index you want to use with the data source connector.")
		kendra_createDataSourceCmd.Flags().String("language-code", "", "The code for a language.")
		kendra_createDataSourceCmd.Flags().String("name", "", "A name for the data source connector.")
		kendra_createDataSourceCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of an IAM role with permission to access the data source and required resources.")
		kendra_createDataSourceCmd.Flags().String("schedule", "", "Sets the frequency for Amazon Kendra to check the documents in your data source repository and update the index.")
		kendra_createDataSourceCmd.Flags().String("tags", "", "A list of key-value pairs that identify or categorize the data source connector.")
		kendra_createDataSourceCmd.Flags().String("type", "", "The type of data source repository.")
		kendra_createDataSourceCmd.Flags().String("vpc-configuration", "", "Configuration information for an Amazon Virtual Private Cloud to connect to your data source.")
		kendra_createDataSourceCmd.MarkFlagRequired("index-id")
		kendra_createDataSourceCmd.MarkFlagRequired("name")
		kendra_createDataSourceCmd.MarkFlagRequired("type")
	})
	kendraCmd.AddCommand(kendra_createDataSourceCmd)
}
