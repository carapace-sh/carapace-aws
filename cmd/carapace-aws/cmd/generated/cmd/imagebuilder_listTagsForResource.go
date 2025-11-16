package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var imagebuilder_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Returns the list of tags for the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imagebuilder_listTagsForResourceCmd).Standalone()

	imagebuilder_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource whose tags you want to retrieve.")
	imagebuilder_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	imagebuilderCmd.AddCommand(imagebuilder_listTagsForResourceCmd)
}
