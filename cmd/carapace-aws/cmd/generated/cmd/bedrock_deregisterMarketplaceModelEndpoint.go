package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_deregisterMarketplaceModelEndpointCmd = &cobra.Command{
	Use:   "deregister-marketplace-model-endpoint",
	Short: "Deregisters an endpoint for a model from Amazon Bedrock Marketplace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_deregisterMarketplaceModelEndpointCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrock_deregisterMarketplaceModelEndpointCmd).Standalone()

		bedrock_deregisterMarketplaceModelEndpointCmd.Flags().String("endpoint-arn", "", "The Amazon Resource Name (ARN) of the endpoint you want to deregister.")
		bedrock_deregisterMarketplaceModelEndpointCmd.MarkFlagRequired("endpoint-arn")
	})
	bedrockCmd.AddCommand(bedrock_deregisterMarketplaceModelEndpointCmd)
}
