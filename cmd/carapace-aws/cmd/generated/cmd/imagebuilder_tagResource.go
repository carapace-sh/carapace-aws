package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var imagebuilder_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds a tag to a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imagebuilder_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(imagebuilder_tagResourceCmd).Standalone()

		imagebuilder_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource that you want to tag.")
		imagebuilder_tagResourceCmd.Flags().String("tags", "", "The tags to apply to the resource.")
		imagebuilder_tagResourceCmd.MarkFlagRequired("resource-arn")
		imagebuilder_tagResourceCmd.MarkFlagRequired("tags")
	})
	imagebuilderCmd.AddCommand(imagebuilder_tagResourceCmd)
}
