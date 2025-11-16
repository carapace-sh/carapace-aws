package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisanalyticsv2_discoverInputSchemaCmd = &cobra.Command{
	Use:   "discover-input-schema",
	Short: "Infers a schema for a SQL-based Kinesis Data Analytics application by evaluating sample records on the specified streaming source (Kinesis data stream or Kinesis Data Firehose delivery stream) or Amazon S3 object.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisanalyticsv2_discoverInputSchemaCmd).Standalone()

	kinesisanalyticsv2_discoverInputSchemaCmd.Flags().String("input-processing-configuration", "", "The [InputProcessingConfiguration]() to use to preprocess the records before discovering the schema of the records.")
	kinesisanalyticsv2_discoverInputSchemaCmd.Flags().String("input-starting-position-configuration", "", "The point at which you want Kinesis Data Analytics to start reading records from the specified streaming source for discovery purposes.")
	kinesisanalyticsv2_discoverInputSchemaCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the streaming source.")
	kinesisanalyticsv2_discoverInputSchemaCmd.Flags().String("s3-configuration", "", "Specify this parameter to discover a schema from data in an Amazon S3 object.")
	kinesisanalyticsv2_discoverInputSchemaCmd.Flags().String("service-execution-role", "", "The ARN of the role that is used to access the streaming source.")
	kinesisanalyticsv2_discoverInputSchemaCmd.MarkFlagRequired("service-execution-role")
	kinesisanalyticsv2Cmd.AddCommand(kinesisanalyticsv2_discoverInputSchemaCmd)
}
