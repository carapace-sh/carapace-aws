package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplifyuibuilder_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Returns a list of tags for a specified Amazon Resource Name (ARN).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplifyuibuilder_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(amplifyuibuilder_listTagsForResourceCmd).Standalone()

		amplifyuibuilder_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) to use to list tags.")
		amplifyuibuilder_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	amplifyuibuilderCmd.AddCommand(amplifyuibuilder_listTagsForResourceCmd)
}
