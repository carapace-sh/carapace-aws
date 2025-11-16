package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_registerMarketplaceModelEndpointCmd = &cobra.Command{
	Use:   "register-marketplace-model-endpoint",
	Short: "Registers an existing Amazon SageMaker endpoint with Amazon Bedrock Marketplace, allowing it to be used with Amazon Bedrock APIs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_registerMarketplaceModelEndpointCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrock_registerMarketplaceModelEndpointCmd).Standalone()

		bedrock_registerMarketplaceModelEndpointCmd.Flags().String("endpoint-identifier", "", "The ARN of the Amazon SageMaker endpoint you want to register with Amazon Bedrock Marketplace.")
		bedrock_registerMarketplaceModelEndpointCmd.Flags().String("model-source-identifier", "", "The ARN of the model from Amazon Bedrock Marketplace that is deployed on the endpoint.")
		bedrock_registerMarketplaceModelEndpointCmd.MarkFlagRequired("endpoint-identifier")
		bedrock_registerMarketplaceModelEndpointCmd.MarkFlagRequired("model-source-identifier")
	})
	bedrockCmd.AddCommand(bedrock_registerMarketplaceModelEndpointCmd)
}
