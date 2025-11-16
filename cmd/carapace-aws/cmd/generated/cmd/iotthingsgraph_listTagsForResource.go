package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotthingsgraph_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists all tags on an AWS IoT Things Graph resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotthingsgraph_listTagsForResourceCmd).Standalone()

	iotthingsgraph_listTagsForResourceCmd.Flags().String("max-results", "", "The maximum number of tags to return.")
	iotthingsgraph_listTagsForResourceCmd.Flags().String("next-token", "", "The token that specifies the next page of results to return.")
	iotthingsgraph_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource whose tags are to be returned.")
	iotthingsgraph_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	iotthingsgraphCmd.AddCommand(iotthingsgraph_listTagsForResourceCmd)
}
