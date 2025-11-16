package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mpa_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Returns a list of the tags for a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mpa_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mpa_listTagsForResourceCmd).Standalone()

		mpa_listTagsForResourceCmd.Flags().String("resource-arn", "", "Amazon Resource Name (ARN) for the resource.")
		mpa_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	mpaCmd.AddCommand(mpa_listTagsForResourceCmd)
}
