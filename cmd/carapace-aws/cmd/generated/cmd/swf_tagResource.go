package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var swf_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Add a tag to a Amazon SWF domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(swf_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(swf_tagResourceCmd).Standalone()

		swf_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) for the Amazon SWF domain.")
		swf_tagResourceCmd.Flags().String("tags", "", "The list of tags to add to a domain.")
		swf_tagResourceCmd.MarkFlagRequired("resource-arn")
		swf_tagResourceCmd.MarkFlagRequired("tags")
	})
	swfCmd.AddCommand(swf_tagResourceCmd)
}
