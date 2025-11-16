package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisanalytics_addApplicationInputCmd = &cobra.Command{
	Use:   "add-application-input",
	Short: "This documentation is for version 1 of the Amazon Kinesis Data Analytics API, which only supports SQL applications.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisanalytics_addApplicationInputCmd).Standalone()

	kinesisanalytics_addApplicationInputCmd.Flags().String("application-name", "", "Name of your existing Amazon Kinesis Analytics application to which you want to add the streaming source.")
	kinesisanalytics_addApplicationInputCmd.Flags().String("current-application-version-id", "", "Current version of your Amazon Kinesis Analytics application.")
	kinesisanalytics_addApplicationInputCmd.Flags().String("input", "", "The [Input](https://docs.aws.amazon.com/kinesisanalytics/latest/dev/API_Input.html) to add.")
	kinesisanalytics_addApplicationInputCmd.MarkFlagRequired("application-name")
	kinesisanalytics_addApplicationInputCmd.MarkFlagRequired("current-application-version-id")
	kinesisanalytics_addApplicationInputCmd.MarkFlagRequired("input")
	kinesisanalyticsCmd.AddCommand(kinesisanalytics_addApplicationInputCmd)
}
