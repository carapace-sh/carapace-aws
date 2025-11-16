package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var textract_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes any tags with the specified keys from the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(textract_untagResourceCmd).Standalone()

	textract_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) that specifies the resource to be untagged.")
	textract_untagResourceCmd.Flags().String("tag-keys", "", "Specifies the tags to be removed from the resource specified by the ResourceARN.")
	textract_untagResourceCmd.MarkFlagRequired("resource-arn")
	textract_untagResourceCmd.MarkFlagRequired("tag-keys")
	textractCmd.AddCommand(textract_untagResourceCmd)
}
