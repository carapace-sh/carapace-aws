package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudhsm_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "This is documentation for **AWS CloudHSM Classic**.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudhsm_listTagsForResourceCmd).Standalone()

	cloudhsm_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the AWS CloudHSM resource.")
	cloudhsm_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	cloudhsmCmd.AddCommand(cloudhsm_listTagsForResourceCmd)
}
