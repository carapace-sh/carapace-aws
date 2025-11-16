package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var marketplaceDeployment_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Tags a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(marketplaceDeployment_tagResourceCmd).Standalone()

	marketplaceDeployment_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) associated with the resource you want to tag.")
	marketplaceDeployment_tagResourceCmd.Flags().String("tags", "", "A map of key-value pairs, where each pair represents a tag present on the resource.")
	marketplaceDeployment_tagResourceCmd.MarkFlagRequired("resource-arn")
	marketplaceDeploymentCmd.AddCommand(marketplaceDeployment_tagResourceCmd)
}
