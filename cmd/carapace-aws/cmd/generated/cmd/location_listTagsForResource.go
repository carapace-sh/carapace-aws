package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var location_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Returns a list of tags that are applied to the specified Amazon Location resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(location_listTagsForResourceCmd).Standalone()

	location_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource whose tags you want to retrieve.")
	location_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	locationCmd.AddCommand(location_listTagsForResourceCmd)
}
