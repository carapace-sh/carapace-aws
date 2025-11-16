package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisanalytics_createApplicationCmd = &cobra.Command{
	Use:   "create-application",
	Short: "This documentation is for version 1 of the Amazon Kinesis Data Analytics API, which only supports SQL applications.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisanalytics_createApplicationCmd).Standalone()

	kinesisanalytics_createApplicationCmd.Flags().String("application-code", "", "One or more SQL statements that read input data, transform it, and generate output.")
	kinesisanalytics_createApplicationCmd.Flags().String("application-description", "", "Summary description of the application.")
	kinesisanalytics_createApplicationCmd.Flags().String("application-name", "", "Name of your Amazon Kinesis Analytics application (for example, `sample-app`).")
	kinesisanalytics_createApplicationCmd.Flags().String("cloud-watch-logging-options", "", "Use this parameter to configure a CloudWatch log stream to monitor application configuration errors.")
	kinesisanalytics_createApplicationCmd.Flags().String("inputs", "", "Use this parameter to configure the application input.")
	kinesisanalytics_createApplicationCmd.Flags().String("outputs", "", "You can configure application output to write data from any of the in-application streams to up to three destinations.")
	kinesisanalytics_createApplicationCmd.Flags().String("tags", "", "A list of one or more tags to assign to the application.")
	kinesisanalytics_createApplicationCmd.MarkFlagRequired("application-name")
	kinesisanalyticsCmd.AddCommand(kinesisanalytics_createApplicationCmd)
}
