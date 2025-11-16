package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector_setTagsForResourceCmd = &cobra.Command{
	Use:   "set-tags-for-resource",
	Short: "Sets tags (key and value pairs) to the assessment template that is specified by the ARN of the assessment template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector_setTagsForResourceCmd).Standalone()

	inspector_setTagsForResourceCmd.Flags().String("resource-arn", "", "The ARN of the assessment template that you want to set tags to.")
	inspector_setTagsForResourceCmd.Flags().String("tags", "", "A collection of key and value pairs that you want to set to the assessment template.")
	inspector_setTagsForResourceCmd.MarkFlagRequired("resource-arn")
	inspectorCmd.AddCommand(inspector_setTagsForResourceCmd)
}
