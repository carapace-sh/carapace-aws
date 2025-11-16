package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var marketplaceDeployment_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists all tags that have been added to a deployment parameter resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(marketplaceDeployment_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(marketplaceDeployment_listTagsForResourceCmd).Standalone()

		marketplaceDeployment_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) associated with the deployment parameter resource you want to list tags on.")
		marketplaceDeployment_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	marketplaceDeploymentCmd.AddCommand(marketplaceDeployment_listTagsForResourceCmd)
}
