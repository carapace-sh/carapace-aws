package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisanalyticsv2_deleteApplicationReferenceDataSourceCmd = &cobra.Command{
	Use:   "delete-application-reference-data-source",
	Short: "Deletes a reference data source configuration from the specified SQL-based Kinesis Data Analytics application's configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisanalyticsv2_deleteApplicationReferenceDataSourceCmd).Standalone()

	kinesisanalyticsv2_deleteApplicationReferenceDataSourceCmd.Flags().String("application-name", "", "The name of an existing application.")
	kinesisanalyticsv2_deleteApplicationReferenceDataSourceCmd.Flags().String("current-application-version-id", "", "The current application version.")
	kinesisanalyticsv2_deleteApplicationReferenceDataSourceCmd.Flags().String("reference-id", "", "The ID of the reference data source.")
	kinesisanalyticsv2_deleteApplicationReferenceDataSourceCmd.MarkFlagRequired("application-name")
	kinesisanalyticsv2_deleteApplicationReferenceDataSourceCmd.MarkFlagRequired("current-application-version-id")
	kinesisanalyticsv2_deleteApplicationReferenceDataSourceCmd.MarkFlagRequired("reference-id")
	kinesisanalyticsv2Cmd.AddCommand(kinesisanalyticsv2_deleteApplicationReferenceDataSourceCmd)
}
