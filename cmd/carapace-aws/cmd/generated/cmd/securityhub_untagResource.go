package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes one or more tags from a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_untagResourceCmd).Standalone()

	securityhub_untagResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource to remove the tags from.")
	securityhub_untagResourceCmd.Flags().String("tag-keys", "", "The tag keys associated with the tags to remove from the resource.")
	securityhub_untagResourceCmd.MarkFlagRequired("resource-arn")
	securityhub_untagResourceCmd.MarkFlagRequired("tag-keys")
	securityhubCmd.AddCommand(securityhub_untagResourceCmd)
}
