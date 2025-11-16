package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_listMarketplaceModelEndpointsCmd = &cobra.Command{
	Use:   "list-marketplace-model-endpoints",
	Short: "Lists the endpoints for models from Amazon Bedrock Marketplace in your Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_listMarketplaceModelEndpointsCmd).Standalone()

	bedrock_listMarketplaceModelEndpointsCmd.Flags().String("max-results", "", "The maximum number of results to return in a single call.")
	bedrock_listMarketplaceModelEndpointsCmd.Flags().String("model-source-equals", "", "If specified, only endpoints for the given model source identifier are returned.")
	bedrock_listMarketplaceModelEndpointsCmd.Flags().String("next-token", "", "The token for the next set of results.")
	bedrockCmd.AddCommand(bedrock_listMarketplaceModelEndpointsCmd)
}
