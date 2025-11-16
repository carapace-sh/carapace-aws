package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53profiles_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists the tags that you associated with the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53profiles_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53profiles_listTagsForResourceCmd).Standalone()

		route53profiles_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) for the resource that you want to list the tags for.")
		route53profiles_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	route53profilesCmd.AddCommand(route53profiles_listTagsForResourceCmd)
}
