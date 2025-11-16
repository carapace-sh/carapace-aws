package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ce_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes one or more tags from a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ce_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ce_untagResourceCmd).Standalone()

		ce_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
		ce_untagResourceCmd.Flags().String("resource-tag-keys", "", "A list of tag keys associated with tags that need to be removed from the resource.")
		ce_untagResourceCmd.MarkFlagRequired("resource-arn")
		ce_untagResourceCmd.MarkFlagRequired("resource-tag-keys")
	})
	ceCmd.AddCommand(ce_untagResourceCmd)
}
