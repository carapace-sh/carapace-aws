package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotthingsgraph_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes a tag from the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotthingsgraph_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotthingsgraph_untagResourceCmd).Standalone()

		iotthingsgraph_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource whose tags are to be removed.")
		iotthingsgraph_untagResourceCmd.Flags().String("tag-keys", "", "A list of tag key names to remove from the resource.")
		iotthingsgraph_untagResourceCmd.MarkFlagRequired("resource-arn")
		iotthingsgraph_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	iotthingsgraphCmd.AddCommand(iotthingsgraph_untagResourceCmd)
}
