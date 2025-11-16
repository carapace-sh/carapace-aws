package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var marketplaceDeployment_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes a tag or list of tags from a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(marketplaceDeployment_untagResourceCmd).Standalone()

	marketplaceDeployment_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) associated with the resource you want to remove the tag from.")
	marketplaceDeployment_untagResourceCmd.Flags().String("tag-keys", "", "A list of key names of tags to be removed.")
	marketplaceDeployment_untagResourceCmd.MarkFlagRequired("resource-arn")
	marketplaceDeployment_untagResourceCmd.MarkFlagRequired("tag-keys")
	marketplaceDeploymentCmd.AddCommand(marketplaceDeployment_untagResourceCmd)
}
