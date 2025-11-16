package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var customerProfiles_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Displays the tags associated with an Amazon Connect Customer Profiles resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(customerProfiles_listTagsForResourceCmd).Standalone()

	customerProfiles_listTagsForResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource for which you want to view tags.")
	customerProfiles_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	customerProfilesCmd.AddCommand(customerProfiles_listTagsForResourceCmd)
}
