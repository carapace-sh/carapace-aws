package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var imagebuilder_getMarketplaceResourceCmd = &cobra.Command{
	Use:   "get-marketplace-resource",
	Short: "Verify the subscription and perform resource dependency checks on the requested Amazon Web Services Marketplace resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imagebuilder_getMarketplaceResourceCmd).Standalone()

	imagebuilder_getMarketplaceResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) that uniquely identifies an Amazon Web Services Marketplace resource.")
	imagebuilder_getMarketplaceResourceCmd.Flags().String("resource-location", "", "The bucket path that you can specify to download the resource from Amazon S3.")
	imagebuilder_getMarketplaceResourceCmd.Flags().String("resource-type", "", "Specifies which type of Amazon Web Services Marketplace resource Image Builder retrieves.")
	imagebuilder_getMarketplaceResourceCmd.MarkFlagRequired("resource-arn")
	imagebuilder_getMarketplaceResourceCmd.MarkFlagRequired("resource-type")
	imagebuilderCmd.AddCommand(imagebuilder_getMarketplaceResourceCmd)
}
