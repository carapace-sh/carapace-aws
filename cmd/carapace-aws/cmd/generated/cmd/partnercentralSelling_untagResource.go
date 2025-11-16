package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var partnercentralSelling_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes a tag or tags from a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(partnercentralSelling_untagResourceCmd).Standalone()

	partnercentralSelling_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource that you want to untag.")
	partnercentralSelling_untagResourceCmd.Flags().String("tag-keys", "", "The keys of the key-value pairs for the tag or tags you want to remove from the specified resource.")
	partnercentralSelling_untagResourceCmd.MarkFlagRequired("resource-arn")
	partnercentralSelling_untagResourceCmd.MarkFlagRequired("tag-keys")
	partnercentralSellingCmd.AddCommand(partnercentralSelling_untagResourceCmd)
}
