package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemakerGeospatial_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "The resource you want to untag.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemakerGeospatial_untagResourceCmd).Standalone()

	sagemakerGeospatial_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource you want to untag.")
	sagemakerGeospatial_untagResourceCmd.Flags().String("tag-keys", "", "Keys of the tags you want to remove.")
	sagemakerGeospatial_untagResourceCmd.MarkFlagRequired("resource-arn")
	sagemakerGeospatial_untagResourceCmd.MarkFlagRequired("tag-keys")
	sagemakerGeospatialCmd.AddCommand(sagemakerGeospatial_untagResourceCmd)
}
