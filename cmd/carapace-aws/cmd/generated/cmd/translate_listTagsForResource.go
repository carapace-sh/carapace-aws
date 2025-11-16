package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var translate_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists all tags associated with a given Amazon Translate resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(translate_listTagsForResourceCmd).Standalone()

	translate_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the given Amazon Translate resource you are querying.")
	translate_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	translateCmd.AddCommand(translate_listTagsForResourceCmd)
}
