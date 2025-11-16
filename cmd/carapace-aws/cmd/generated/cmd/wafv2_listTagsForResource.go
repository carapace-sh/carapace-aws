package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafv2_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Retrieves the [TagInfoForResource]() for the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafv2_listTagsForResourceCmd).Standalone()

	wafv2_listTagsForResourceCmd.Flags().String("limit", "", "The maximum number of objects that you want WAF to return for this request.")
	wafv2_listTagsForResourceCmd.Flags().String("next-marker", "", "When you request a list of objects with a `Limit` setting, if the number of objects that are still available for retrieval exceeds the limit, WAF returns a `NextMarker` value in the response.")
	wafv2_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
	wafv2_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	wafv2Cmd.AddCommand(wafv2_listTagsForResourceCmd)
}
