package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var licenseManagerLinuxSubscriptions_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Remove one or more metadata tag from the specified Amazon Web Services resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(licenseManagerLinuxSubscriptions_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(licenseManagerLinuxSubscriptions_untagResourceCmd).Standalone()

		licenseManagerLinuxSubscriptions_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the Amazon Web Services resource to remove the metadata tags from.")
		licenseManagerLinuxSubscriptions_untagResourceCmd.Flags().String("tag-keys", "", "A list of metadata tag keys to remove from the requested resource.")
		licenseManagerLinuxSubscriptions_untagResourceCmd.MarkFlagRequired("resource-arn")
		licenseManagerLinuxSubscriptions_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	licenseManagerLinuxSubscriptionsCmd.AddCommand(licenseManagerLinuxSubscriptions_untagResourceCmd)
}
