package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qbusiness_updateDataSourceCmd = &cobra.Command{
	Use:   "update-data-source",
	Short: "Updates an existing Amazon Q Business data source connector.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qbusiness_updateDataSourceCmd).Standalone()

	qbusiness_updateDataSourceCmd.Flags().String("application-id", "", "The identifier of the Amazon Q Business application the data source is attached to.")
	qbusiness_updateDataSourceCmd.Flags().String("configuration", "", "")
	qbusiness_updateDataSourceCmd.Flags().String("data-source-id", "", "The identifier of the data source connector.")
	qbusiness_updateDataSourceCmd.Flags().String("description", "", "The description of the data source connector.")
	qbusiness_updateDataSourceCmd.Flags().String("display-name", "", "A name of the data source connector.")
	qbusiness_updateDataSourceCmd.Flags().String("document-enrichment-configuration", "", "")
	qbusiness_updateDataSourceCmd.Flags().String("index-id", "", "The identifier of the index attached to the data source connector.")
	qbusiness_updateDataSourceCmd.Flags().String("media-extraction-configuration", "", "The configuration for extracting information from media in documents for your data source.")
	qbusiness_updateDataSourceCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of an IAM role with permission to access the data source and required resources.")
	qbusiness_updateDataSourceCmd.Flags().String("sync-schedule", "", "The chosen update frequency for your data source.")
	qbusiness_updateDataSourceCmd.Flags().String("vpc-configuration", "", "")
	qbusiness_updateDataSourceCmd.MarkFlagRequired("application-id")
	qbusiness_updateDataSourceCmd.MarkFlagRequired("data-source-id")
	qbusiness_updateDataSourceCmd.MarkFlagRequired("index-id")
	qbusinessCmd.AddCommand(qbusiness_updateDataSourceCmd)
}
