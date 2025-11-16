package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Retrieves a list of resource tags for a resource arn.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_listTagsForResourceCmd).Standalone()

	greengrass_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
	greengrass_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	greengrassCmd.AddCommand(greengrass_listTagsForResourceCmd)
}
