package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var finspace_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "A list of all tags for a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(finspace_listTagsForResourceCmd).Standalone()

	finspace_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name of the resource.")
	finspace_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	finspaceCmd.AddCommand(finspace_listTagsForResourceCmd)
}
