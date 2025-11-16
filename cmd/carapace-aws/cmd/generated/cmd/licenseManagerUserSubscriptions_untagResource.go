package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var licenseManagerUserSubscriptions_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes tags from a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(licenseManagerUserSubscriptions_untagResourceCmd).Standalone()

	licenseManagerUserSubscriptions_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource that you want to remove tags from.")
	licenseManagerUserSubscriptions_untagResourceCmd.Flags().String("tag-keys", "", "The tag keys to remove from the resource.")
	licenseManagerUserSubscriptions_untagResourceCmd.MarkFlagRequired("resource-arn")
	licenseManagerUserSubscriptions_untagResourceCmd.MarkFlagRequired("tag-keys")
	licenseManagerUserSubscriptionsCmd.AddCommand(licenseManagerUserSubscriptions_untagResourceCmd)
}
