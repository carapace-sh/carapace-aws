package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lambda_updateEventSourceMappingCmd = &cobra.Command{
	Use:   "update-event-source-mapping",
	Short: "Updates an event source mapping.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lambda_updateEventSourceMappingCmd).Standalone()

	lambda_updateEventSourceMappingCmd.Flags().String("amazon-managed-kafka-event-source-config", "", "")
	lambda_updateEventSourceMappingCmd.Flags().String("batch-size", "", "The maximum number of records in each batch that Lambda pulls from your stream or queue and sends to your function.")
	lambda_updateEventSourceMappingCmd.Flags().String("bisect-batch-on-function-error", "", "(Kinesis and DynamoDB Streams only) If the function returns an error, split the batch in two and retry.")
	lambda_updateEventSourceMappingCmd.Flags().String("destination-config", "", "(Kinesis, DynamoDB Streams, Amazon MSK, and self-managed Kafka only) A configuration object that specifies the destination of an event after Lambda processes it.")
	lambda_updateEventSourceMappingCmd.Flags().String("document-dbevent-source-config", "", "Specific configuration settings for a DocumentDB event source.")
	lambda_updateEventSourceMappingCmd.Flags().String("enabled", "", "When true, the event source mapping is active.")
	lambda_updateEventSourceMappingCmd.Flags().String("filter-criteria", "", "An object that defines the filter criteria that determine whether Lambda should process an event.")
	lambda_updateEventSourceMappingCmd.Flags().String("function-name", "", "The name or ARN of the Lambda function.")
	lambda_updateEventSourceMappingCmd.Flags().String("function-response-types", "", "(Kinesis, DynamoDB Streams, and Amazon SQS) A list of current response type enums applied to the event source mapping.")
	lambda_updateEventSourceMappingCmd.Flags().String("kmskey-arn", "", "The ARN of the Key Management Service (KMS) customer managed key that Lambda uses to encrypt your function's [filter criteria](https://docs.aws.amazon.com/lambda/latest/dg/invocation-eventfiltering.html#filtering-basics). By default, Lambda does not encrypt your filter criteria object.")
	lambda_updateEventSourceMappingCmd.Flags().String("maximum-batching-window-in-seconds", "", "The maximum amount of time, in seconds, that Lambda spends gathering records before invoking the function.")
	lambda_updateEventSourceMappingCmd.Flags().String("maximum-record-age-in-seconds", "", "(Kinesis and DynamoDB Streams only) Discard records older than the specified age.")
	lambda_updateEventSourceMappingCmd.Flags().String("maximum-retry-attempts", "", "(Kinesis and DynamoDB Streams only) Discard records after the specified number of retries.")
	lambda_updateEventSourceMappingCmd.Flags().String("metrics-config", "", "The metrics configuration for your event source.")
	lambda_updateEventSourceMappingCmd.Flags().String("parallelization-factor", "", "(Kinesis and DynamoDB Streams only) The number of batches to process from each shard concurrently.")
	lambda_updateEventSourceMappingCmd.Flags().String("provisioned-poller-config", "", "(Amazon MSK and self-managed Apache Kafka only) The provisioned mode configuration for the event source.")
	lambda_updateEventSourceMappingCmd.Flags().String("scaling-config", "", "(Amazon SQS only) The scaling configuration for the event source.")
	lambda_updateEventSourceMappingCmd.Flags().String("self-managed-kafka-event-source-config", "", "")
	lambda_updateEventSourceMappingCmd.Flags().String("source-access-configurations", "", "An array of authentication protocols or VPC components required to secure your event source.")
	lambda_updateEventSourceMappingCmd.Flags().String("tumbling-window-in-seconds", "", "(Kinesis and DynamoDB Streams only) The duration in seconds of a processing window for DynamoDB and Kinesis Streams event sources.")
	lambda_updateEventSourceMappingCmd.Flags().String("uuid", "", "The identifier of the event source mapping.")
	lambda_updateEventSourceMappingCmd.MarkFlagRequired("uuid")
	lambdaCmd.AddCommand(lambda_updateEventSourceMappingCmd)
}
