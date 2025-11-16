package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector2_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes tags from a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector2_untagResourceCmd).Standalone()

	inspector2_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) for the resource to remove tags from.")
	inspector2_untagResourceCmd.Flags().String("tag-keys", "", "The tag keys to remove from the resource.")
	inspector2_untagResourceCmd.MarkFlagRequired("resource-arn")
	inspector2_untagResourceCmd.MarkFlagRequired("tag-keys")
	inspector2Cmd.AddCommand(inspector2_untagResourceCmd)
}
