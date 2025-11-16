package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kms_listResourceTagsCmd = &cobra.Command{
	Use:   "list-resource-tags",
	Short: "Returns all tags on the specified KMS key.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kms_listResourceTagsCmd).Standalone()

	kms_listResourceTagsCmd.Flags().String("key-id", "", "Gets tags on the specified KMS key.")
	kms_listResourceTagsCmd.Flags().String("limit", "", "Use this parameter to specify the maximum number of items to return.")
	kms_listResourceTagsCmd.Flags().String("marker", "", "Use this parameter in a subsequent request after you receive a response with truncated results.")
	kms_listResourceTagsCmd.MarkFlagRequired("key-id")
	kmsCmd.AddCommand(kms_listResourceTagsCmd)
}
