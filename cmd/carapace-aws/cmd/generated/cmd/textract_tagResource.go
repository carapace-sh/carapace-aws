package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var textract_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds one or more tags to the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(textract_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(textract_tagResourceCmd).Standalone()

		textract_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) that specifies the resource to be tagged.")
		textract_tagResourceCmd.Flags().String("tags", "", "A set of tags (key-value pairs) that you want to assign to the resource.")
		textract_tagResourceCmd.MarkFlagRequired("resource-arn")
		textract_tagResourceCmd.MarkFlagRequired("tags")
	})
	textractCmd.AddCommand(textract_tagResourceCmd)
}
