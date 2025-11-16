package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisanalytics_addApplicationReferenceDataSourceCmd = &cobra.Command{
	Use:   "add-application-reference-data-source",
	Short: "This documentation is for version 1 of the Amazon Kinesis Data Analytics API, which only supports SQL applications.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisanalytics_addApplicationReferenceDataSourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kinesisanalytics_addApplicationReferenceDataSourceCmd).Standalone()

		kinesisanalytics_addApplicationReferenceDataSourceCmd.Flags().String("application-name", "", "Name of an existing application.")
		kinesisanalytics_addApplicationReferenceDataSourceCmd.Flags().String("current-application-version-id", "", "Version of the application for which you are adding the reference data source.")
		kinesisanalytics_addApplicationReferenceDataSourceCmd.Flags().String("reference-data-source", "", "The reference data source can be an object in your Amazon S3 bucket.")
		kinesisanalytics_addApplicationReferenceDataSourceCmd.MarkFlagRequired("application-name")
		kinesisanalytics_addApplicationReferenceDataSourceCmd.MarkFlagRequired("current-application-version-id")
		kinesisanalytics_addApplicationReferenceDataSourceCmd.MarkFlagRequired("reference-data-source")
	})
	kinesisanalyticsCmd.AddCommand(kinesisanalytics_addApplicationReferenceDataSourceCmd)
}
