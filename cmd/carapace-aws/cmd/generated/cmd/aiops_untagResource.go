package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var aiops_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes one or more tags from the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(aiops_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(aiops_untagResourceCmd).Standalone()

		aiops_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource that you want to remove the tags from.")
		aiops_untagResourceCmd.Flags().String("tag-keys", "", "The list of tag keys to remove from the resource.")
		aiops_untagResourceCmd.MarkFlagRequired("resource-arn")
		aiops_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	aiopsCmd.AddCommand(aiops_untagResourceCmd)
}
