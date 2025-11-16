package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transfer_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists all of the tags associated with the Amazon Resource Name (ARN) that you specify.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transfer_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(transfer_listTagsForResourceCmd).Standalone()

		transfer_listTagsForResourceCmd.Flags().String("arn", "", "Requests the tags associated with a particular Amazon Resource Name (ARN).")
		transfer_listTagsForResourceCmd.Flags().String("max-results", "", "Specifies the number of tags to return as a response to the `ListTagsForResource` request.")
		transfer_listTagsForResourceCmd.Flags().String("next-token", "", "When you request additional results from the `ListTagsForResource` operation, a `NextToken` parameter is returned in the input.")
		transfer_listTagsForResourceCmd.MarkFlagRequired("arn")
	})
	transferCmd.AddCommand(transfer_listTagsForResourceCmd)
}
