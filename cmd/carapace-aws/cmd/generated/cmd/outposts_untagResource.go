package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var outposts_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes tags from the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(outposts_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(outposts_untagResourceCmd).Standalone()

		outposts_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
		outposts_untagResourceCmd.Flags().String("tag-keys", "", "The tag keys.")
		outposts_untagResourceCmd.MarkFlagRequired("resource-arn")
		outposts_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	outpostsCmd.AddCommand(outposts_untagResourceCmd)
}
