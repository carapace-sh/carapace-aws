package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudhsm_addTagsToResourceCmd = &cobra.Command{
	Use:   "add-tags-to-resource",
	Short: "This is documentation for **AWS CloudHSM Classic**.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudhsm_addTagsToResourceCmd).Standalone()

	cloudhsm_addTagsToResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the AWS CloudHSM resource to tag.")
	cloudhsm_addTagsToResourceCmd.Flags().String("tag-list", "", "One or more tags.")
	cloudhsm_addTagsToResourceCmd.MarkFlagRequired("resource-arn")
	cloudhsm_addTagsToResourceCmd.MarkFlagRequired("tag-list")
	cloudhsmCmd.AddCommand(cloudhsm_addTagsToResourceCmd)
}
