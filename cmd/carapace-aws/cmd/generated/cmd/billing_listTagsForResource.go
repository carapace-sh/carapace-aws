package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var billing_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists tags associated with the billing view resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(billing_listTagsForResourceCmd).Standalone()

	billing_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
	billing_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	billingCmd.AddCommand(billing_listTagsForResourceCmd)
}
