package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kendra_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds the specified tag to the specified index, FAQ, data source, or other resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kendra_tagResourceCmd).Standalone()

	kendra_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the index, FAQ, data source, or other resource to add a tag.")
	kendra_tagResourceCmd.Flags().String("tags", "", "A list of tag keys to add to the index, FAQ, data source, or other resource.")
	kendra_tagResourceCmd.MarkFlagRequired("resource-arn")
	kendra_tagResourceCmd.MarkFlagRequired("tags")
	kendraCmd.AddCommand(kendra_tagResourceCmd)
}
