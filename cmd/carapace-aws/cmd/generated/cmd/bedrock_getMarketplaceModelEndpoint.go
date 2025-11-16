package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_getMarketplaceModelEndpointCmd = &cobra.Command{
	Use:   "get-marketplace-model-endpoint",
	Short: "Retrieves details about a specific endpoint for a model from Amazon Bedrock Marketplace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_getMarketplaceModelEndpointCmd).Standalone()

	bedrock_getMarketplaceModelEndpointCmd.Flags().String("endpoint-arn", "", "The Amazon Resource Name (ARN) of the endpoint you want to get information about.")
	bedrock_getMarketplaceModelEndpointCmd.MarkFlagRequired("endpoint-arn")
	bedrockCmd.AddCommand(bedrock_getMarketplaceModelEndpointCmd)
}
