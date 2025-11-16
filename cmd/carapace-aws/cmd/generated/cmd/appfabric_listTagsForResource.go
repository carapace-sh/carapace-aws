package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appfabric_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Returns a list of tags for a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appfabric_listTagsForResourceCmd).Standalone()

	appfabric_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource for which you want to retrieve tags.")
	appfabric_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	appfabricCmd.AddCommand(appfabric_listTagsForResourceCmd)
}
