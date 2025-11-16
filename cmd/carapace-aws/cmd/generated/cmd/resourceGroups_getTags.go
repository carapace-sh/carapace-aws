package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resourceGroups_getTagsCmd = &cobra.Command{
	Use:   "get-tags",
	Short: "Returns a list of tags that are associated with a resource group, specified by an Amazon resource name (ARN).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resourceGroups_getTagsCmd).Standalone()

	resourceGroups_getTagsCmd.Flags().String("arn", "", "The Amazon resource name (ARN) of the resource group whose tags you want to retrieve.")
	resourceGroups_getTagsCmd.MarkFlagRequired("arn")
	resourceGroupsCmd.AddCommand(resourceGroups_getTagsCmd)
}
