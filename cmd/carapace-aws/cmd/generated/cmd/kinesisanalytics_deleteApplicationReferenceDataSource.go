package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisanalytics_deleteApplicationReferenceDataSourceCmd = &cobra.Command{
	Use:   "delete-application-reference-data-source",
	Short: "This documentation is for version 1 of the Amazon Kinesis Data Analytics API, which only supports SQL applications.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisanalytics_deleteApplicationReferenceDataSourceCmd).Standalone()

	kinesisanalytics_deleteApplicationReferenceDataSourceCmd.Flags().String("application-name", "", "Name of an existing application.")
	kinesisanalytics_deleteApplicationReferenceDataSourceCmd.Flags().String("current-application-version-id", "", "Version of the application.")
	kinesisanalytics_deleteApplicationReferenceDataSourceCmd.Flags().String("reference-id", "", "ID of the reference data source.")
	kinesisanalytics_deleteApplicationReferenceDataSourceCmd.MarkFlagRequired("application-name")
	kinesisanalytics_deleteApplicationReferenceDataSourceCmd.MarkFlagRequired("current-application-version-id")
	kinesisanalytics_deleteApplicationReferenceDataSourceCmd.MarkFlagRequired("reference-id")
	kinesisanalyticsCmd.AddCommand(kinesisanalytics_deleteApplicationReferenceDataSourceCmd)
}
