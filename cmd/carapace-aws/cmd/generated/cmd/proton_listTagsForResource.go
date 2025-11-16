package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proton_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "List tags for a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proton_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(proton_listTagsForResourceCmd).Standalone()

		proton_listTagsForResourceCmd.Flags().String("max-results", "", "The maximum number of tags to list.")
		proton_listTagsForResourceCmd.Flags().String("next-token", "", "A token that indicates the location of the next resource tag in the array of resource tags, after the list of resource tags that was previously requested.")
		proton_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource for the listed tags.")
		proton_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	protonCmd.AddCommand(proton_listTagsForResourceCmd)
}
