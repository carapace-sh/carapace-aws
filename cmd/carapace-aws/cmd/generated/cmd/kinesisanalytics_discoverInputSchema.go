package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisanalytics_discoverInputSchemaCmd = &cobra.Command{
	Use:   "discover-input-schema",
	Short: "This documentation is for version 1 of the Amazon Kinesis Data Analytics API, which only supports SQL applications.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisanalytics_discoverInputSchemaCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kinesisanalytics_discoverInputSchemaCmd).Standalone()

		kinesisanalytics_discoverInputSchemaCmd.Flags().String("input-processing-configuration", "", "The [InputProcessingConfiguration](https://docs.aws.amazon.com/kinesisanalytics/latest/dev/API_InputProcessingConfiguration.html) to use to preprocess the records before discovering the schema of the records.")
		kinesisanalytics_discoverInputSchemaCmd.Flags().String("input-starting-position-configuration", "", "Point at which you want Amazon Kinesis Analytics to start reading records from the specified streaming source discovery purposes.")
		kinesisanalytics_discoverInputSchemaCmd.Flags().String("resource-arn", "", "Amazon Resource Name (ARN) of the streaming source.")
		kinesisanalytics_discoverInputSchemaCmd.Flags().String("role-arn", "", "ARN of the IAM role that Amazon Kinesis Analytics can assume to access the stream on your behalf.")
		kinesisanalytics_discoverInputSchemaCmd.Flags().String("s3-configuration", "", "Specify this parameter to discover a schema from data in an Amazon S3 object.")
	})
	kinesisanalyticsCmd.AddCommand(kinesisanalytics_discoverInputSchemaCmd)
}
