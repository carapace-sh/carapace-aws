package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rtbfabric_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists tags for a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rtbfabric_listTagsForResourceCmd).Standalone()

	rtbfabric_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource for which you want to retrieve tags.")
	rtbfabric_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	rtbfabricCmd.AddCommand(rtbfabric_listTagsForResourceCmd)
}
