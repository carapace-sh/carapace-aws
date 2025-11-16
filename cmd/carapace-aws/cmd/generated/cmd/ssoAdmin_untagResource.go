package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssoAdmin_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Disassociates a set of tags from a specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssoAdmin_untagResourceCmd).Standalone()

	ssoAdmin_untagResourceCmd.Flags().String("instance-arn", "", "The ARN of the IAM Identity Center instance under which the operation will be executed.")
	ssoAdmin_untagResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource with the tags to be listed.")
	ssoAdmin_untagResourceCmd.Flags().String("tag-keys", "", "The keys of tags that are attached to the resource.")
	ssoAdmin_untagResourceCmd.MarkFlagRequired("resource-arn")
	ssoAdmin_untagResourceCmd.MarkFlagRequired("tag-keys")
	ssoAdminCmd.AddCommand(ssoAdmin_untagResourceCmd)
}
