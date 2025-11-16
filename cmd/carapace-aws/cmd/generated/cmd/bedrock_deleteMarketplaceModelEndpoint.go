package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_deleteMarketplaceModelEndpointCmd = &cobra.Command{
	Use:   "delete-marketplace-model-endpoint",
	Short: "Deletes an endpoint for a model from Amazon Bedrock Marketplace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_deleteMarketplaceModelEndpointCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrock_deleteMarketplaceModelEndpointCmd).Standalone()

		bedrock_deleteMarketplaceModelEndpointCmd.Flags().String("endpoint-arn", "", "The Amazon Resource Name (ARN) of the endpoint you want to delete.")
		bedrock_deleteMarketplaceModelEndpointCmd.MarkFlagRequired("endpoint-arn")
	})
	bedrockCmd.AddCommand(bedrock_deleteMarketplaceModelEndpointCmd)
}
