package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emrServerless_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists the tags assigned to the resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emrServerless_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(emrServerless_listTagsForResourceCmd).Standalone()

		emrServerless_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) that identifies the resource to list the tags for.")
		emrServerless_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	emrServerlessCmd.AddCommand(emrServerless_listTagsForResourceCmd)
}
