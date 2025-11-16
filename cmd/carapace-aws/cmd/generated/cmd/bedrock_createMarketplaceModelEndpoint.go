package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_createMarketplaceModelEndpointCmd = &cobra.Command{
	Use:   "create-marketplace-model-endpoint",
	Short: "Creates an endpoint for a model from Amazon Bedrock Marketplace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_createMarketplaceModelEndpointCmd).Standalone()

	bedrock_createMarketplaceModelEndpointCmd.Flags().String("accept-eula", "", "Indicates whether you accept the end-user license agreement (EULA) for the model.")
	bedrock_createMarketplaceModelEndpointCmd.Flags().String("client-request-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	bedrock_createMarketplaceModelEndpointCmd.Flags().String("endpoint-config", "", "The configuration for the endpoint, including the number and type of instances to use.")
	bedrock_createMarketplaceModelEndpointCmd.Flags().String("endpoint-name", "", "The name of the endpoint.")
	bedrock_createMarketplaceModelEndpointCmd.Flags().String("model-source-identifier", "", "The ARN of the model from Amazon Bedrock Marketplace that you want to deploy to the endpoint.")
	bedrock_createMarketplaceModelEndpointCmd.Flags().String("tags", "", "An array of key-value pairs to apply to the underlying Amazon SageMaker endpoint.")
	bedrock_createMarketplaceModelEndpointCmd.MarkFlagRequired("endpoint-config")
	bedrock_createMarketplaceModelEndpointCmd.MarkFlagRequired("endpoint-name")
	bedrock_createMarketplaceModelEndpointCmd.MarkFlagRequired("model-source-identifier")
	bedrockCmd.AddCommand(bedrock_createMarketplaceModelEndpointCmd)
}
