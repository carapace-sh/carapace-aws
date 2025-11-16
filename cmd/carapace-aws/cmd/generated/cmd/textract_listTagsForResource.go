package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var textract_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists all tags for an Amazon Textract resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(textract_listTagsForResourceCmd).Standalone()

	textract_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) that specifies the resource to list tags for.")
	textract_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	textractCmd.AddCommand(textract_listTagsForResourceCmd)
}
