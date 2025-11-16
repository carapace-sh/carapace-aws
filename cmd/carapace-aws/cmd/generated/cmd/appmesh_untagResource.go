package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appmesh_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Deletes specified tags from a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appmesh_untagResourceCmd).Standalone()

	appmesh_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource to delete tags from.")
	appmesh_untagResourceCmd.Flags().String("tag-keys", "", "The keys of the tags to be removed.")
	appmesh_untagResourceCmd.MarkFlagRequired("resource-arn")
	appmesh_untagResourceCmd.MarkFlagRequired("tag-keys")
	appmeshCmd.AddCommand(appmesh_untagResourceCmd)
}
