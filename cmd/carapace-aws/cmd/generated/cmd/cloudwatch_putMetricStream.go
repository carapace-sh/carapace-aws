package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudwatch_putMetricStreamCmd = &cobra.Command{
	Use:   "put-metric-stream",
	Short: "Creates or updates a metric stream.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudwatch_putMetricStreamCmd).Standalone()

	cloudwatch_putMetricStreamCmd.Flags().String("exclude-filters", "", "If you specify this parameter, the stream sends metrics from all metric namespaces except for the namespaces that you specify here.")
	cloudwatch_putMetricStreamCmd.Flags().String("firehose-arn", "", "The ARN of the Amazon Kinesis Data Firehose delivery stream to use for this metric stream.")
	cloudwatch_putMetricStreamCmd.Flags().String("include-filters", "", "If you specify this parameter, the stream sends only the metrics from the metric namespaces that you specify here.")
	cloudwatch_putMetricStreamCmd.Flags().String("include-linked-accounts-metrics", "", "If you are creating a metric stream in a monitoring account, specify `true` to include metrics from source accounts in the metric stream.")
	cloudwatch_putMetricStreamCmd.Flags().String("name", "", "If you are creating a new metric stream, this is the name for the new stream.")
	cloudwatch_putMetricStreamCmd.Flags().String("output-format", "", "The output format for the stream.")
	cloudwatch_putMetricStreamCmd.Flags().String("role-arn", "", "The ARN of an IAM role that this metric stream will use to access Amazon Kinesis Data Firehose resources.")
	cloudwatch_putMetricStreamCmd.Flags().String("statistics-configurations", "", "By default, a metric stream always sends the `MAX`, `MIN`, `SUM`, and `SAMPLECOUNT` statistics for each metric that is streamed.")
	cloudwatch_putMetricStreamCmd.Flags().String("tags", "", "A list of key-value pairs to associate with the metric stream.")
	cloudwatch_putMetricStreamCmd.MarkFlagRequired("firehose-arn")
	cloudwatch_putMetricStreamCmd.MarkFlagRequired("name")
	cloudwatch_putMetricStreamCmd.MarkFlagRequired("output-format")
	cloudwatch_putMetricStreamCmd.MarkFlagRequired("role-arn")
	cloudwatchCmd.AddCommand(cloudwatch_putMetricStreamCmd)
}
