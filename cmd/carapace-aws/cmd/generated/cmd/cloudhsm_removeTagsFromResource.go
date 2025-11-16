package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudhsm_removeTagsFromResourceCmd = &cobra.Command{
	Use:   "remove-tags-from-resource",
	Short: "This is documentation for **AWS CloudHSM Classic**.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudhsm_removeTagsFromResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudhsm_removeTagsFromResourceCmd).Standalone()

		cloudhsm_removeTagsFromResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the AWS CloudHSM resource.")
		cloudhsm_removeTagsFromResourceCmd.Flags().String("tag-key-list", "", "The tag key or keys to remove.")
		cloudhsm_removeTagsFromResourceCmd.MarkFlagRequired("resource-arn")
		cloudhsm_removeTagsFromResourceCmd.MarkFlagRequired("tag-key-list")
	})
	cloudhsmCmd.AddCommand(cloudhsm_removeTagsFromResourceCmd)
}
