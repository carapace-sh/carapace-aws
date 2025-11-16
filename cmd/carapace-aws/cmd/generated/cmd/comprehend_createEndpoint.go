package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehend_createEndpointCmd = &cobra.Command{
	Use:   "create-endpoint",
	Short: "Creates a model-specific endpoint for synchronous inference for a previously trained custom model For information about endpoints, see [Managing endpoints](https://docs.aws.amazon.com/comprehend/latest/dg/manage-endpoints.html).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehend_createEndpointCmd).Standalone()

	comprehend_createEndpointCmd.Flags().String("client-request-token", "", "An idempotency token provided by the customer.")
	comprehend_createEndpointCmd.Flags().String("data-access-role-arn", "", "The Amazon Resource Name (ARN) of the IAM role that grants Amazon Comprehend read access to trained custom models encrypted with a customer managed key (ModelKmsKeyId).")
	comprehend_createEndpointCmd.Flags().String("desired-inference-units", "", "The desired number of inference units to be used by the model using this endpoint.")
	comprehend_createEndpointCmd.Flags().String("endpoint-name", "", "This is the descriptive suffix that becomes part of the `EndpointArn` used for all subsequent requests to this resource.")
	comprehend_createEndpointCmd.Flags().String("flywheel-arn", "", "The Amazon Resource Number (ARN) of the flywheel to which the endpoint will be attached.")
	comprehend_createEndpointCmd.Flags().String("model-arn", "", "The Amazon Resource Number (ARN) of the model to which the endpoint will be attached.")
	comprehend_createEndpointCmd.Flags().String("tags", "", "Tags to associate with the endpoint.")
	comprehend_createEndpointCmd.MarkFlagRequired("desired-inference-units")
	comprehend_createEndpointCmd.MarkFlagRequired("endpoint-name")
	comprehendCmd.AddCommand(comprehend_createEndpointCmd)
}
