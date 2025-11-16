package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rtbfabric_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes a tag or tags from a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rtbfabric_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rtbfabric_untagResourceCmd).Standalone()

		rtbfabric_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource that you want to untag.")
		rtbfabric_untagResourceCmd.Flags().String("tag-keys", "", "The keys of the key-value pairs for the tag or tags you want to remove from the specified resource.")
		rtbfabric_untagResourceCmd.MarkFlagRequired("resource-arn")
		rtbfabric_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	rtbfabricCmd.AddCommand(rtbfabric_untagResourceCmd)
}
