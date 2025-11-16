package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var clouddirectory_listPolicyAttachmentsCmd = &cobra.Command{
	Use:   "list-policy-attachments",
	Short: "Returns all of the `ObjectIdentifiers` to which a given policy is attached.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clouddirectory_listPolicyAttachmentsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(clouddirectory_listPolicyAttachmentsCmd).Standalone()

		clouddirectory_listPolicyAttachmentsCmd.Flags().String("consistency-level", "", "Represents the manner and timing in which the successful write or update of an object is reflected in a subsequent read operation of that same object.")
		clouddirectory_listPolicyAttachmentsCmd.Flags().String("directory-arn", "", "The Amazon Resource Name (ARN) that is associated with the [Directory]() where objects reside.")
		clouddirectory_listPolicyAttachmentsCmd.Flags().String("max-results", "", "The maximum number of items to be retrieved in a single call.")
		clouddirectory_listPolicyAttachmentsCmd.Flags().String("next-token", "", "The pagination token.")
		clouddirectory_listPolicyAttachmentsCmd.Flags().String("policy-reference", "", "The reference that identifies the policy object.")
		clouddirectory_listPolicyAttachmentsCmd.MarkFlagRequired("directory-arn")
		clouddirectory_listPolicyAttachmentsCmd.MarkFlagRequired("policy-reference")
	})
	clouddirectoryCmd.AddCommand(clouddirectory_listPolicyAttachmentsCmd)
}
