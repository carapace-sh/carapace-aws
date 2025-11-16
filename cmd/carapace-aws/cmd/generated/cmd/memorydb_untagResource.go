package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var memorydb_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Use this operation to remove tags on a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(memorydb_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(memorydb_untagResourceCmd).Standalone()

		memorydb_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource to which the tags are to be removed.")
		memorydb_untagResourceCmd.Flags().String("tag-keys", "", "The list of keys of the tags that are to be removed.")
		memorydb_untagResourceCmd.MarkFlagRequired("resource-arn")
		memorydb_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	memorydbCmd.AddCommand(memorydb_untagResourceCmd)
}
