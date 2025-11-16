package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var billing_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes one or more tags from a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(billing_untagResourceCmd).Standalone()

	billing_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
	billing_untagResourceCmd.Flags().String("resource-tag-keys", "", "A list of tag key value pairs that are associated with the resource.")
	billing_untagResourceCmd.MarkFlagRequired("resource-arn")
	billing_untagResourceCmd.MarkFlagRequired("resource-tag-keys")
	billingCmd.AddCommand(billing_untagResourceCmd)
}
