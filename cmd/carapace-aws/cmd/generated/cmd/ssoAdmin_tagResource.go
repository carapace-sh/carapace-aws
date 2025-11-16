package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssoAdmin_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Associates a set of tags with a specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssoAdmin_tagResourceCmd).Standalone()

	ssoAdmin_tagResourceCmd.Flags().String("instance-arn", "", "The ARN of the IAM Identity Center instance under which the operation will be executed.")
	ssoAdmin_tagResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource with the tags to be listed.")
	ssoAdmin_tagResourceCmd.Flags().String("tags", "", "A set of key-value pairs that are used to manage the resource.")
	ssoAdmin_tagResourceCmd.MarkFlagRequired("resource-arn")
	ssoAdmin_tagResourceCmd.MarkFlagRequired("tags")
	ssoAdminCmd.AddCommand(ssoAdmin_tagResourceCmd)
}
