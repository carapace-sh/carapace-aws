package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var customerProfiles_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes one or more tags from the specified Amazon Connect Customer Profiles resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(customerProfiles_untagResourceCmd).Standalone()

	customerProfiles_untagResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource from which you are removing tags.")
	customerProfiles_untagResourceCmd.Flags().String("tag-keys", "", "The list of tag keys to remove from the resource.")
	customerProfiles_untagResourceCmd.MarkFlagRequired("resource-arn")
	customerProfiles_untagResourceCmd.MarkFlagRequired("tag-keys")
	customerProfilesCmd.AddCommand(customerProfiles_untagResourceCmd)
}
