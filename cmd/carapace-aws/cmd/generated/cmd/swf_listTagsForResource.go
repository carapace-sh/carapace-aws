package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var swf_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "List tags for a given domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(swf_listTagsForResourceCmd).Standalone()

	swf_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) for the Amazon SWF domain.")
	swf_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	swfCmd.AddCommand(swf_listTagsForResourceCmd)
}
