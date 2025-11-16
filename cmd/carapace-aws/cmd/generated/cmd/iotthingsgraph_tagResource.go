package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotthingsgraph_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Creates a tag for the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotthingsgraph_tagResourceCmd).Standalone()

	iotthingsgraph_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource whose tags are returned.")
	iotthingsgraph_tagResourceCmd.Flags().String("tags", "", "A list of tags to add to the resource.&gt;")
	iotthingsgraph_tagResourceCmd.MarkFlagRequired("resource-arn")
	iotthingsgraph_tagResourceCmd.MarkFlagRequired("tags")
	iotthingsgraphCmd.AddCommand(iotthingsgraph_tagResourceCmd)
}
