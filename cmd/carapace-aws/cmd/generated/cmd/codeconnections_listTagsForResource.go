package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeconnections_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Gets the set of key-value pairs (metadata) that are used to manage the resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeconnections_listTagsForResourceCmd).Standalone()

	codeconnections_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource for which you want to get information about tags, if any.")
	codeconnections_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	codeconnectionsCmd.AddCommand(codeconnections_listTagsForResourceCmd)
}
