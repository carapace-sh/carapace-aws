package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appmesh_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "List the tags for an App Mesh resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appmesh_listTagsForResourceCmd).Standalone()

	appmesh_listTagsForResourceCmd.Flags().String("limit", "", "The maximum number of tag results returned by `ListTagsForResource` in paginated output.")
	appmesh_listTagsForResourceCmd.Flags().String("next-token", "", "The `nextToken` value returned from a previous paginated `ListTagsForResource` request where `limit` was used and the results exceeded the value of that parameter.")
	appmesh_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) that identifies the resource to list the tags for.")
	appmesh_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	appmeshCmd.AddCommand(appmesh_listTagsForResourceCmd)
}
