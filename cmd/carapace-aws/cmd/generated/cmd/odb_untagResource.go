package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var odb_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes tags from the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(odb_untagResourceCmd).Standalone()

	odb_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource to remove tags from.")
	odb_untagResourceCmd.Flags().String("tag-keys", "", "The names (keys) of the tags to remove from the resource.")
	odb_untagResourceCmd.MarkFlagRequired("resource-arn")
	odb_untagResourceCmd.MarkFlagRequired("tag-keys")
	odbCmd.AddCommand(odb_untagResourceCmd)
}
