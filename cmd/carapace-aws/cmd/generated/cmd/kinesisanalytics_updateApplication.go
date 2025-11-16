package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisanalytics_updateApplicationCmd = &cobra.Command{
	Use:   "update-application",
	Short: "This documentation is for version 1 of the Amazon Kinesis Data Analytics API, which only supports SQL applications.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisanalytics_updateApplicationCmd).Standalone()

	kinesisanalytics_updateApplicationCmd.Flags().String("application-name", "", "Name of the Amazon Kinesis Analytics application to update.")
	kinesisanalytics_updateApplicationCmd.Flags().String("application-update", "", "Describes application updates.")
	kinesisanalytics_updateApplicationCmd.Flags().String("current-application-version-id", "", "The current application version ID.")
	kinesisanalytics_updateApplicationCmd.MarkFlagRequired("application-name")
	kinesisanalytics_updateApplicationCmd.MarkFlagRequired("application-update")
	kinesisanalytics_updateApplicationCmd.MarkFlagRequired("current-application-version-id")
	kinesisanalyticsCmd.AddCommand(kinesisanalytics_updateApplicationCmd)
}
