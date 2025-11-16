package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemakerRuntime_invokeEndpointAsyncCmd = &cobra.Command{
	Use:   "invoke-endpoint-async",
	Short: "After you deploy a model into production using Amazon SageMaker AI hosting services, your client applications use this API to get inferences from the model hosted at the specified endpoint in an asynchronous manner.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemakerRuntime_invokeEndpointAsyncCmd).Standalone()

	sagemakerRuntime_invokeEndpointAsyncCmd.Flags().String("accept", "", "The desired MIME type of the inference response from the model container.")
	sagemakerRuntime_invokeEndpointAsyncCmd.Flags().String("content-type", "", "The MIME type of the input data in the request body.")
	sagemakerRuntime_invokeEndpointAsyncCmd.Flags().String("custom-attributes", "", "Provides additional information about a request for an inference submitted to a model hosted at an Amazon SageMaker AI endpoint.")
	sagemakerRuntime_invokeEndpointAsyncCmd.Flags().String("endpoint-name", "", "The name of the endpoint that you specified when you created the endpoint using the [CreateEndpoint](https://docs.aws.amazon.com/sagemaker/latest/dg/API_CreateEndpoint.html) API.")
	sagemakerRuntime_invokeEndpointAsyncCmd.Flags().String("inference-id", "", "The identifier for the inference request.")
	sagemakerRuntime_invokeEndpointAsyncCmd.Flags().String("input-location", "", "The Amazon S3 URI where the inference request payload is stored.")
	sagemakerRuntime_invokeEndpointAsyncCmd.Flags().String("invocation-timeout-seconds", "", "Maximum amount of time in seconds a request can be processed before it is marked as expired.")
	sagemakerRuntime_invokeEndpointAsyncCmd.Flags().String("request-ttlseconds", "", "Maximum age in seconds a request can be in the queue before it is marked as expired.")
	sagemakerRuntime_invokeEndpointAsyncCmd.MarkFlagRequired("endpoint-name")
	sagemakerRuntime_invokeEndpointAsyncCmd.MarkFlagRequired("input-location")
	sagemakerRuntimeCmd.AddCommand(sagemakerRuntime_invokeEndpointAsyncCmd)
}
