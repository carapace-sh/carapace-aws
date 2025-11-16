package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssoAdmin_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists the tags that are attached to a specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssoAdmin_listTagsForResourceCmd).Standalone()

	ssoAdmin_listTagsForResourceCmd.Flags().String("instance-arn", "", "The ARN of the IAM Identity Center instance under which the operation will be executed.")
	ssoAdmin_listTagsForResourceCmd.Flags().String("next-token", "", "The pagination token for the list API.")
	ssoAdmin_listTagsForResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource with the tags to be listed.")
	ssoAdmin_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	ssoAdminCmd.AddCommand(ssoAdmin_listTagsForResourceCmd)
}
