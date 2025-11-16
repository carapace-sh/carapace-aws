package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var invoicing_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists the tags for a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(invoicing_listTagsForResourceCmd).Standalone()

	invoicing_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of tags to list.")
	invoicing_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	invoicingCmd.AddCommand(invoicing_listTagsForResourceCmd)
}
