package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var billing_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "An API operation for adding one or more tags (key-value pairs) to a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(billing_tagResourceCmd).Standalone()

	billing_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
	billing_tagResourceCmd.Flags().String("resource-tags", "", "A list of tag key value pairs that are associated with the resource.")
	billing_tagResourceCmd.MarkFlagRequired("resource-arn")
	billing_tagResourceCmd.MarkFlagRequired("resource-tags")
	billingCmd.AddCommand(billing_tagResourceCmd)
}
