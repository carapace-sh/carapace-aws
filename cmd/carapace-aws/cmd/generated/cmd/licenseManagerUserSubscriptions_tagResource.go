package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var licenseManagerUserSubscriptions_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds tags to a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(licenseManagerUserSubscriptions_tagResourceCmd).Standalone()

	licenseManagerUserSubscriptions_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource that you want to tag.")
	licenseManagerUserSubscriptions_tagResourceCmd.Flags().String("tags", "", "The tags to apply to the specified resource.")
	licenseManagerUserSubscriptions_tagResourceCmd.MarkFlagRequired("resource-arn")
	licenseManagerUserSubscriptions_tagResourceCmd.MarkFlagRequired("tags")
	licenseManagerUserSubscriptionsCmd.AddCommand(licenseManagerUserSubscriptions_tagResourceCmd)
}
