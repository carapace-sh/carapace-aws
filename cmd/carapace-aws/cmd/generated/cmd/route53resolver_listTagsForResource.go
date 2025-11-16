package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53resolver_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists the tags that you associated with the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53resolver_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53resolver_listTagsForResourceCmd).Standalone()

		route53resolver_listTagsForResourceCmd.Flags().String("max-results", "", "The maximum number of tags that you want to return in the response to a `ListTagsForResource` request.")
		route53resolver_listTagsForResourceCmd.Flags().String("next-token", "", "For the first `ListTagsForResource` request, omit this value.")
		route53resolver_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) for the resource that you want to list tags for.")
		route53resolver_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	route53resolverCmd.AddCommand(route53resolver_listTagsForResourceCmd)
}
