package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisanalytics_addApplicationOutputCmd = &cobra.Command{
	Use:   "add-application-output",
	Short: "This documentation is for version 1 of the Amazon Kinesis Data Analytics API, which only supports SQL applications.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisanalytics_addApplicationOutputCmd).Standalone()

	kinesisanalytics_addApplicationOutputCmd.Flags().String("application-name", "", "Name of the application to which you want to add the output configuration.")
	kinesisanalytics_addApplicationOutputCmd.Flags().String("current-application-version-id", "", "Version of the application to which you want to add the output configuration.")
	kinesisanalytics_addApplicationOutputCmd.Flags().String("output", "", "An array of objects, each describing one output configuration.")
	kinesisanalytics_addApplicationOutputCmd.MarkFlagRequired("application-name")
	kinesisanalytics_addApplicationOutputCmd.MarkFlagRequired("current-application-version-id")
	kinesisanalytics_addApplicationOutputCmd.MarkFlagRequired("output")
	kinesisanalyticsCmd.AddCommand(kinesisanalytics_addApplicationOutputCmd)
}
