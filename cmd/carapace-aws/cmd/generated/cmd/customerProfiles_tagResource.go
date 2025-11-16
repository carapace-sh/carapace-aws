package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var customerProfiles_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Assigns one or more tags (key-value pairs) to the specified Amazon Connect Customer Profiles resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(customerProfiles_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(customerProfiles_tagResourceCmd).Standalone()

		customerProfiles_tagResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource that you're adding tags to.")
		customerProfiles_tagResourceCmd.Flags().String("tags", "", "The tags used to organize, track, or control access for this resource.")
		customerProfiles_tagResourceCmd.MarkFlagRequired("resource-arn")
		customerProfiles_tagResourceCmd.MarkFlagRequired("tags")
	})
	customerProfilesCmd.AddCommand(customerProfiles_tagResourceCmd)
}
