package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var personalize_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Get a list of [tags](https://docs.aws.amazon.com/personalize/latest/dg/tagging-resources.html) attached to a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(personalize_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(personalize_listTagsForResourceCmd).Standalone()

		personalize_listTagsForResourceCmd.Flags().String("resource-arn", "", "The resource's Amazon Resource Name (ARN).")
		personalize_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	personalizeCmd.AddCommand(personalize_listTagsForResourceCmd)
}
