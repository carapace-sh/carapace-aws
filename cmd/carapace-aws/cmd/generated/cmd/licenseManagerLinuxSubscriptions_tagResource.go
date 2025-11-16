package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var licenseManagerLinuxSubscriptions_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Add metadata tags to the specified Amazon Web Services resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(licenseManagerLinuxSubscriptions_tagResourceCmd).Standalone()

	licenseManagerLinuxSubscriptions_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the Amazon Web Services resource to which to add the specified metadata tags.")
	licenseManagerLinuxSubscriptions_tagResourceCmd.Flags().String("tags", "", "The metadata tags to assign to the Amazon Web Services resource.")
	licenseManagerLinuxSubscriptions_tagResourceCmd.MarkFlagRequired("resource-arn")
	licenseManagerLinuxSubscriptions_tagResourceCmd.MarkFlagRequired("tags")
	licenseManagerLinuxSubscriptionsCmd.AddCommand(licenseManagerLinuxSubscriptions_tagResourceCmd)
}
