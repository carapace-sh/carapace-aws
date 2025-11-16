package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivs_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes tags from the resource with the specified ARN.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivs_untagResourceCmd).Standalone()

	ivs_untagResourceCmd.Flags().String("resource-arn", "", "ARN of the resource for which tags are to be removed.")
	ivs_untagResourceCmd.Flags().String("tag-keys", "", "Array of tags to be removed.")
	ivs_untagResourceCmd.MarkFlagRequired("resource-arn")
	ivs_untagResourceCmd.MarkFlagRequired("tag-keys")
	ivsCmd.AddCommand(ivs_untagResourceCmd)
}
