package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lambda_createEventSourceMappingCmd = &cobra.Command{
	Use:   "create-event-source-mapping",
	Short: "Creates a mapping between an event source and an Lambda function.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lambda_createEventSourceMappingCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lambda_createEventSourceMappingCmd).Standalone()

		lambda_createEventSourceMappingCmd.Flags().String("amazon-managed-kafka-event-source-config", "", "Specific configuration settings for an Amazon Managed Streaming for Apache Kafka (Amazon MSK) event source.")
		lambda_createEventSourceMappingCmd.Flags().String("batch-size", "", "The maximum number of records in each batch that Lambda pulls from your stream or queue and sends to your function.")
		lambda_createEventSourceMappingCmd.Flags().String("bisect-batch-on-function-error", "", "(Kinesis and DynamoDB Streams only) If the function returns an error, split the batch in two and retry.")
		lambda_createEventSourceMappingCmd.Flags().String("destination-config", "", "(Kinesis, DynamoDB Streams, Amazon MSK, and self-managed Kafka only) A configuration object that specifies the destination of an event after Lambda processes it.")
		lambda_createEventSourceMappingCmd.Flags().String("document-dbevent-source-config", "", "Specific configuration settings for a DocumentDB event source.")
		lambda_createEventSourceMappingCmd.Flags().String("enabled", "", "When true, the event source mapping is active.")
		lambda_createEventSourceMappingCmd.Flags().String("event-source-arn", "", "The Amazon Resource Name (ARN) of the event source.")
		lambda_createEventSourceMappingCmd.Flags().String("filter-criteria", "", "An object that defines the filter criteria that determine whether Lambda should process an event.")
		lambda_createEventSourceMappingCmd.Flags().String("function-name", "", "The name or ARN of the Lambda function.")
		lambda_createEventSourceMappingCmd.Flags().String("function-response-types", "", "(Kinesis, DynamoDB Streams, and Amazon SQS) A list of current response type enums applied to the event source mapping.")
		lambda_createEventSourceMappingCmd.Flags().String("kmskey-arn", "", "The ARN of the Key Management Service (KMS) customer managed key that Lambda uses to encrypt your function's [filter criteria](https://docs.aws.amazon.com/lambda/latest/dg/invocation-eventfiltering.html#filtering-basics). By default, Lambda does not encrypt your filter criteria object.")
		lambda_createEventSourceMappingCmd.Flags().String("maximum-batching-window-in-seconds", "", "The maximum amount of time, in seconds, that Lambda spends gathering records before invoking the function.")
		lambda_createEventSourceMappingCmd.Flags().String("maximum-record-age-in-seconds", "", "(Kinesis and DynamoDB Streams only) Discard records older than the specified age.")
		lambda_createEventSourceMappingCmd.Flags().String("maximum-retry-attempts", "", "(Kinesis and DynamoDB Streams only) Discard records after the specified number of retries.")
		lambda_createEventSourceMappingCmd.Flags().String("metrics-config", "", "The metrics configuration for your event source.")
		lambda_createEventSourceMappingCmd.Flags().String("parallelization-factor", "", "(Kinesis and DynamoDB Streams only) The number of batches to process from each shard concurrently.")
		lambda_createEventSourceMappingCmd.Flags().String("provisioned-poller-config", "", "(Amazon MSK and self-managed Apache Kafka only) The provisioned mode configuration for the event source.")
		lambda_createEventSourceMappingCmd.Flags().String("queues", "", "(MQ) The name of the Amazon MQ broker destination queue to consume.")
		lambda_createEventSourceMappingCmd.Flags().String("scaling-config", "", "(Amazon SQS only) The scaling configuration for the event source.")
		lambda_createEventSourceMappingCmd.Flags().String("self-managed-event-source", "", "The self-managed Apache Kafka cluster to receive records from.")
		lambda_createEventSourceMappingCmd.Flags().String("self-managed-kafka-event-source-config", "", "Specific configuration settings for a self-managed Apache Kafka event source.")
		lambda_createEventSourceMappingCmd.Flags().String("source-access-configurations", "", "An array of authentication protocols or VPC components required to secure your event source.")
		lambda_createEventSourceMappingCmd.Flags().String("starting-position", "", "The position in a stream from which to start reading.")
		lambda_createEventSourceMappingCmd.Flags().String("starting-position-timestamp", "", "With `StartingPosition` set to `AT_TIMESTAMP`, the time from which to start reading.")
		lambda_createEventSourceMappingCmd.Flags().String("tags", "", "A list of tags to apply to the event source mapping.")
		lambda_createEventSourceMappingCmd.Flags().String("topics", "", "The name of the Kafka topic.")
		lambda_createEventSourceMappingCmd.Flags().String("tumbling-window-in-seconds", "", "(Kinesis and DynamoDB Streams only) The duration in seconds of a processing window for DynamoDB and Kinesis Streams event sources.")
		lambda_createEventSourceMappingCmd.MarkFlagRequired("function-name")
	})
	lambdaCmd.AddCommand(lambda_createEventSourceMappingCmd)
}
