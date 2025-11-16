package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisanalytics_stopApplicationCmd = &cobra.Command{
	Use:   "stop-application",
	Short: "This documentation is for version 1 of the Amazon Kinesis Data Analytics API, which only supports SQL applications.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisanalytics_stopApplicationCmd).Standalone()

	kinesisanalytics_stopApplicationCmd.Flags().String("application-name", "", "Name of the running application to stop.")
	kinesisanalytics_stopApplicationCmd.MarkFlagRequired("application-name")
	kinesisanalyticsCmd.AddCommand(kinesisanalytics_stopApplicationCmd)
}
