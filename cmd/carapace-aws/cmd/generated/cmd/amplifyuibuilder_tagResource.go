package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplifyuibuilder_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Tags the resource with a tag key and value.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplifyuibuilder_tagResourceCmd).Standalone()

	amplifyuibuilder_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) to use to tag a resource.")
	amplifyuibuilder_tagResourceCmd.Flags().String("tags", "", "A list of tag key value pairs for a specified Amazon Resource Name (ARN).")
	amplifyuibuilder_tagResourceCmd.MarkFlagRequired("resource-arn")
	amplifyuibuilder_tagResourceCmd.MarkFlagRequired("tags")
	amplifyuibuilderCmd.AddCommand(amplifyuibuilder_tagResourceCmd)
}
