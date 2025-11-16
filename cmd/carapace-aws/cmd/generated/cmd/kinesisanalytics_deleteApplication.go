package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisanalytics_deleteApplicationCmd = &cobra.Command{
	Use:   "delete-application",
	Short: "This documentation is for version 1 of the Amazon Kinesis Data Analytics API, which only supports SQL applications.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisanalytics_deleteApplicationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kinesisanalytics_deleteApplicationCmd).Standalone()

		kinesisanalytics_deleteApplicationCmd.Flags().String("application-name", "", "Name of the Amazon Kinesis Analytics application to delete.")
		kinesisanalytics_deleteApplicationCmd.Flags().String("create-timestamp", "", "You can use the `DescribeApplication` operation to get this value.")
		kinesisanalytics_deleteApplicationCmd.MarkFlagRequired("application-name")
		kinesisanalytics_deleteApplicationCmd.MarkFlagRequired("create-timestamp")
	})
	kinesisanalyticsCmd.AddCommand(kinesisanalytics_deleteApplicationCmd)
}
