package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplify_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Returns a list of tags for a specified Amazon Resource Name (ARN).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplify_listTagsForResourceCmd).Standalone()

	amplify_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) to use to list tags.")
	amplify_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	amplifyCmd.AddCommand(amplify_listTagsForResourceCmd)
}
