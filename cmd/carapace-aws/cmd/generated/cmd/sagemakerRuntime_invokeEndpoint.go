package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemakerRuntime_invokeEndpointCmd = &cobra.Command{
	Use:   "invoke-endpoint",
	Short: "After you deploy a model into production using Amazon SageMaker AI hosting services, your client applications use this API to get inferences from the model hosted at the specified endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemakerRuntime_invokeEndpointCmd).Standalone()

	sagemakerRuntime_invokeEndpointCmd.Flags().String("accept", "", "The desired MIME type of the inference response from the model container.")
	sagemakerRuntime_invokeEndpointCmd.Flags().String("body", "", "Provides input data, in the format specified in the `ContentType` request header.")
	sagemakerRuntime_invokeEndpointCmd.Flags().String("content-type", "", "The MIME type of the input data in the request body.")
	sagemakerRuntime_invokeEndpointCmd.Flags().String("custom-attributes", "", "Provides additional information about a request for an inference submitted to a model hosted at an Amazon SageMaker AI endpoint.")
	sagemakerRuntime_invokeEndpointCmd.Flags().String("enable-explanations", "", "An optional JMESPath expression used to override the `EnableExplanations` parameter of the `ClarifyExplainerConfig` API.")
	sagemakerRuntime_invokeEndpointCmd.Flags().String("endpoint-name", "", "The name of the endpoint that you specified when you created the endpoint using the [CreateEndpoint](https://docs.aws.amazon.com/sagemaker/latest/dg/API_CreateEndpoint.html) API.")
	sagemakerRuntime_invokeEndpointCmd.Flags().String("inference-component-name", "", "If the endpoint hosts one or more inference components, this parameter specifies the name of inference component to invoke.")
	sagemakerRuntime_invokeEndpointCmd.Flags().String("inference-id", "", "If you provide a value, it is added to the captured data when you enable data capture on the endpoint.")
	sagemakerRuntime_invokeEndpointCmd.Flags().String("session-id", "", "Creates a stateful session or identifies an existing one.")
	sagemakerRuntime_invokeEndpointCmd.Flags().String("target-container-hostname", "", "If the endpoint hosts multiple containers and is configured to use direct invocation, this parameter specifies the host name of the container to invoke.")
	sagemakerRuntime_invokeEndpointCmd.Flags().String("target-model", "", "The model to request for inference when invoking a multi-model endpoint.")
	sagemakerRuntime_invokeEndpointCmd.Flags().String("target-variant", "", "Specify the production variant to send the inference request to when invoking an endpoint that is running two or more variants.")
	sagemakerRuntime_invokeEndpointCmd.MarkFlagRequired("body")
	sagemakerRuntime_invokeEndpointCmd.MarkFlagRequired("endpoint-name")
	sagemakerRuntimeCmd.AddCommand(sagemakerRuntime_invokeEndpointCmd)
}
