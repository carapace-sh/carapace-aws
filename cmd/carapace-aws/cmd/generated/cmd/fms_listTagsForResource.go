package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fms_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Retrieves the list of tags for the specified Amazon Web Services resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fms_listTagsForResourceCmd).Standalone()

	fms_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource to return tags for.")
	fms_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	fmsCmd.AddCommand(fms_listTagsForResourceCmd)
}
