package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes tags from a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_untagResourceCmd).Standalone()

		glue_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource from which to remove the tags.")
		glue_untagResourceCmd.Flags().String("tags-to-remove", "", "Tags to remove from this resource.")
		glue_untagResourceCmd.MarkFlagRequired("resource-arn")
		glue_untagResourceCmd.MarkFlagRequired("tags-to-remove")
	})
	glueCmd.AddCommand(glue_untagResourceCmd)
}
