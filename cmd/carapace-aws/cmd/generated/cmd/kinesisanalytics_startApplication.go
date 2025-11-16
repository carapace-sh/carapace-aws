package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisanalytics_startApplicationCmd = &cobra.Command{
	Use:   "start-application",
	Short: "This documentation is for version 1 of the Amazon Kinesis Data Analytics API, which only supports SQL applications.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisanalytics_startApplicationCmd).Standalone()

	kinesisanalytics_startApplicationCmd.Flags().String("application-name", "", "Name of the application.")
	kinesisanalytics_startApplicationCmd.Flags().String("input-configurations", "", "Identifies the specific input, by ID, that the application starts consuming.")
	kinesisanalytics_startApplicationCmd.MarkFlagRequired("application-name")
	kinesisanalytics_startApplicationCmd.MarkFlagRequired("input-configurations")
	kinesisanalyticsCmd.AddCommand(kinesisanalytics_startApplicationCmd)
}
