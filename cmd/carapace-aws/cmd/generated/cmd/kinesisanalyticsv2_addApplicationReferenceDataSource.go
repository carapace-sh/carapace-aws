package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisanalyticsv2_addApplicationReferenceDataSourceCmd = &cobra.Command{
	Use:   "add-application-reference-data-source",
	Short: "Adds a reference data source to an existing SQL-based Kinesis Data Analytics application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisanalyticsv2_addApplicationReferenceDataSourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kinesisanalyticsv2_addApplicationReferenceDataSourceCmd).Standalone()

		kinesisanalyticsv2_addApplicationReferenceDataSourceCmd.Flags().String("application-name", "", "The name of an existing application.")
		kinesisanalyticsv2_addApplicationReferenceDataSourceCmd.Flags().String("current-application-version-id", "", "The version of the application for which you are adding the reference data source.")
		kinesisanalyticsv2_addApplicationReferenceDataSourceCmd.Flags().String("reference-data-source", "", "The reference data source can be an object in your Amazon S3 bucket.")
		kinesisanalyticsv2_addApplicationReferenceDataSourceCmd.MarkFlagRequired("application-name")
		kinesisanalyticsv2_addApplicationReferenceDataSourceCmd.MarkFlagRequired("current-application-version-id")
		kinesisanalyticsv2_addApplicationReferenceDataSourceCmd.MarkFlagRequired("reference-data-source")
	})
	kinesisanalyticsv2Cmd.AddCommand(kinesisanalyticsv2_addApplicationReferenceDataSourceCmd)
}
