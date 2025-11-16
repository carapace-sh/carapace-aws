package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_removeTagsFromResourceCmd = &cobra.Command{
	Use:   "remove-tags-from-resource",
	Short: "Removes tag keys from the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_removeTagsFromResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssm_removeTagsFromResourceCmd).Standalone()

		ssm_removeTagsFromResourceCmd.Flags().String("resource-id", "", "The ID of the resource from which you want to remove tags.")
		ssm_removeTagsFromResourceCmd.Flags().String("resource-type", "", "The type of resource from which you want to remove a tag.")
		ssm_removeTagsFromResourceCmd.Flags().String("tag-keys", "", "Tag keys that you want to remove from the specified resource.")
		ssm_removeTagsFromResourceCmd.MarkFlagRequired("resource-id")
		ssm_removeTagsFromResourceCmd.MarkFlagRequired("resource-type")
		ssm_removeTagsFromResourceCmd.MarkFlagRequired("tag-keys")
	})
	ssmCmd.AddCommand(ssm_removeTagsFromResourceCmd)
}
