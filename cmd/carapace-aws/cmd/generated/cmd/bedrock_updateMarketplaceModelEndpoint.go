package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_updateMarketplaceModelEndpointCmd = &cobra.Command{
	Use:   "update-marketplace-model-endpoint",
	Short: "Updates the configuration of an existing endpoint for a model from Amazon Bedrock Marketplace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_updateMarketplaceModelEndpointCmd).Standalone()

	bedrock_updateMarketplaceModelEndpointCmd.Flags().String("client-request-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	bedrock_updateMarketplaceModelEndpointCmd.Flags().String("endpoint-arn", "", "The Amazon Resource Name (ARN) of the endpoint you want to update.")
	bedrock_updateMarketplaceModelEndpointCmd.Flags().String("endpoint-config", "", "The new configuration for the endpoint, including the number and type of instances to use.")
	bedrock_updateMarketplaceModelEndpointCmd.MarkFlagRequired("endpoint-arn")
	bedrock_updateMarketplaceModelEndpointCmd.MarkFlagRequired("endpoint-config")
	bedrockCmd.AddCommand(bedrock_updateMarketplaceModelEndpointCmd)
}
