package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var location_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes one or more tags from the specified Amazon Location resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(location_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(location_untagResourceCmd).Standalone()

		location_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource from which you want to remove tags.")
		location_untagResourceCmd.Flags().String("tag-keys", "", "The list of tag keys to remove from the specified resource.")
		location_untagResourceCmd.MarkFlagRequired("resource-arn")
		location_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	locationCmd.AddCommand(location_untagResourceCmd)
}
