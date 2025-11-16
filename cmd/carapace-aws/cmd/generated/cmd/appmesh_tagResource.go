package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appmesh_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Associates the specified tags to a resource with the specified `resourceArn`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appmesh_tagResourceCmd).Standalone()

	appmesh_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource to add tags to.")
	appmesh_tagResourceCmd.Flags().String("tags", "", "The tags to add to the resource.")
	appmesh_tagResourceCmd.MarkFlagRequired("resource-arn")
	appmesh_tagResourceCmd.MarkFlagRequired("tags")
	appmeshCmd.AddCommand(appmesh_tagResourceCmd)
}
