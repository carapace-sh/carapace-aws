package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var partnercentralSelling_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Returns a list of tags for a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(partnercentralSelling_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(partnercentralSelling_listTagsForResourceCmd).Standalone()

		partnercentralSelling_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource for which you want to retrieve tags.")
		partnercentralSelling_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	partnercentralSellingCmd.AddCommand(partnercentralSelling_listTagsForResourceCmd)
}
