package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemakerRuntime_invokeEndpointWithResponseStreamCmd = &cobra.Command{
	Use:   "invoke-endpoint-with-response-stream",
	Short: "Invokes a model at the specified endpoint to return the inference response as a stream.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemakerRuntime_invokeEndpointWithResponseStreamCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemakerRuntime_invokeEndpointWithResponseStreamCmd).Standalone()

		sagemakerRuntime_invokeEndpointWithResponseStreamCmd.Flags().String("accept", "", "The desired MIME type of the inference response from the model container.")
		sagemakerRuntime_invokeEndpointWithResponseStreamCmd.Flags().String("body", "", "Provides input data, in the format specified in the `ContentType` request header.")
		sagemakerRuntime_invokeEndpointWithResponseStreamCmd.Flags().String("content-type", "", "The MIME type of the input data in the request body.")
		sagemakerRuntime_invokeEndpointWithResponseStreamCmd.Flags().String("custom-attributes", "", "Provides additional information about a request for an inference submitted to a model hosted at an Amazon SageMaker AI endpoint.")
		sagemakerRuntime_invokeEndpointWithResponseStreamCmd.Flags().String("endpoint-name", "", "The name of the endpoint that you specified when you created the endpoint using the [CreateEndpoint](https://docs.aws.amazon.com/sagemaker/latest/dg/API_CreateEndpoint.html) API.")
		sagemakerRuntime_invokeEndpointWithResponseStreamCmd.Flags().String("inference-component-name", "", "If the endpoint hosts one or more inference components, this parameter specifies the name of inference component to invoke for a streaming response.")
		sagemakerRuntime_invokeEndpointWithResponseStreamCmd.Flags().String("inference-id", "", "An identifier that you assign to your request.")
		sagemakerRuntime_invokeEndpointWithResponseStreamCmd.Flags().String("session-id", "", "The ID of a stateful session to handle your request.")
		sagemakerRuntime_invokeEndpointWithResponseStreamCmd.Flags().String("target-container-hostname", "", "If the endpoint hosts multiple containers and is configured to use direct invocation, this parameter specifies the host name of the container to invoke.")
		sagemakerRuntime_invokeEndpointWithResponseStreamCmd.Flags().String("target-variant", "", "Specify the production variant to send the inference request to when invoking an endpoint that is running two or more variants.")
		sagemakerRuntime_invokeEndpointWithResponseStreamCmd.MarkFlagRequired("body")
		sagemakerRuntime_invokeEndpointWithResponseStreamCmd.MarkFlagRequired("endpoint-name")
	})
	sagemakerRuntimeCmd.AddCommand(sagemakerRuntime_invokeEndpointWithResponseStreamCmd)
}
