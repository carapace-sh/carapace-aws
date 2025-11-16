package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kendra_updateDataSourceCmd = &cobra.Command{
	Use:   "update-data-source",
	Short: "Updates an Amazon Kendra data source connector.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kendra_updateDataSourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kendra_updateDataSourceCmd).Standalone()

		kendra_updateDataSourceCmd.Flags().String("configuration", "", "Configuration information you want to update for the data source connector.")
		kendra_updateDataSourceCmd.Flags().String("custom-document-enrichment-configuration", "", "Configuration information you want to update for altering document metadata and content during the document ingestion process.")
		kendra_updateDataSourceCmd.Flags().String("description", "", "A new description for the data source connector.")
		kendra_updateDataSourceCmd.Flags().String("id", "", "The identifier of the data source connector you want to update.")
		kendra_updateDataSourceCmd.Flags().String("index-id", "", "The identifier of the index used with the data source connector.")
		kendra_updateDataSourceCmd.Flags().String("language-code", "", "The code for a language you want to update for the data source connector.")
		kendra_updateDataSourceCmd.Flags().String("name", "", "A new name for the data source connector.")
		kendra_updateDataSourceCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of an IAM role with permission to access the data source and required resources.")
		kendra_updateDataSourceCmd.Flags().String("schedule", "", "The sync schedule you want to update for the data source connector.")
		kendra_updateDataSourceCmd.Flags().String("vpc-configuration", "", "Configuration information for an Amazon Virtual Private Cloud to connect to your data source.")
		kendra_updateDataSourceCmd.MarkFlagRequired("id")
		kendra_updateDataSourceCmd.MarkFlagRequired("index-id")
	})
	kendraCmd.AddCommand(kendra_updateDataSourceCmd)
}
