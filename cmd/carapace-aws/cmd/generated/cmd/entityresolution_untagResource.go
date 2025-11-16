package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var entityresolution_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes one or more tags from the specified Entity Resolution resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(entityresolution_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(entityresolution_untagResourceCmd).Standalone()

		entityresolution_untagResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource for which you want to untag.")
		entityresolution_untagResourceCmd.Flags().String("tag-keys", "", "The list of tag keys to remove from the resource.")
		entityresolution_untagResourceCmd.MarkFlagRequired("resource-arn")
		entityresolution_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	entityresolutionCmd.AddCommand(entityresolution_untagResourceCmd)
}
