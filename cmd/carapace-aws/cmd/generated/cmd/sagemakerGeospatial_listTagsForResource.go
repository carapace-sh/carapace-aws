package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemakerGeospatial_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists the tags attached to the resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemakerGeospatial_listTagsForResourceCmd).Standalone()

	sagemakerGeospatial_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource you want to tag.")
	sagemakerGeospatial_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	sagemakerGeospatialCmd.AddCommand(sagemakerGeospatial_listTagsForResourceCmd)
}
