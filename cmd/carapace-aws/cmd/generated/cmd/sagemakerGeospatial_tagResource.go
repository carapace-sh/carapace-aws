package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemakerGeospatial_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "The resource you want to tag.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemakerGeospatial_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemakerGeospatial_tagResourceCmd).Standalone()

		sagemakerGeospatial_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource you want to tag.")
		sagemakerGeospatial_tagResourceCmd.Flags().String("tags", "", "Each tag consists of a key and a value.")
		sagemakerGeospatial_tagResourceCmd.MarkFlagRequired("resource-arn")
		sagemakerGeospatial_tagResourceCmd.MarkFlagRequired("tags")
	})
	sagemakerGeospatialCmd.AddCommand(sagemakerGeospatial_tagResourceCmd)
}
