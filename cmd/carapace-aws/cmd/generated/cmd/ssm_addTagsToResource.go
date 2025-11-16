package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_addTagsToResourceCmd = &cobra.Command{
	Use:   "add-tags-to-resource",
	Short: "Adds or overwrites one or more tags for the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_addTagsToResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssm_addTagsToResourceCmd).Standalone()

		ssm_addTagsToResourceCmd.Flags().String("resource-id", "", "The resource ID you want to tag.")
		ssm_addTagsToResourceCmd.Flags().String("resource-type", "", "Specifies the type of resource you are tagging.")
		ssm_addTagsToResourceCmd.Flags().String("tags", "", "One or more tags.")
		ssm_addTagsToResourceCmd.MarkFlagRequired("resource-id")
		ssm_addTagsToResourceCmd.MarkFlagRequired("resource-type")
		ssm_addTagsToResourceCmd.MarkFlagRequired("tags")
	})
	ssmCmd.AddCommand(ssm_addTagsToResourceCmd)
}
