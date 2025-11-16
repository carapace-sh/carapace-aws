package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var clouddirectory_listObjectPoliciesCmd = &cobra.Command{
	Use:   "list-object-policies",
	Short: "Returns policies attached to an object in pagination fashion.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clouddirectory_listObjectPoliciesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(clouddirectory_listObjectPoliciesCmd).Standalone()

		clouddirectory_listObjectPoliciesCmd.Flags().String("consistency-level", "", "Represents the manner and timing in which the successful write or update of an object is reflected in a subsequent read operation of that same object.")
		clouddirectory_listObjectPoliciesCmd.Flags().String("directory-arn", "", "The Amazon Resource Name (ARN) that is associated with the [Directory]() where objects reside.")
		clouddirectory_listObjectPoliciesCmd.Flags().String("max-results", "", "The maximum number of items to be retrieved in a single call.")
		clouddirectory_listObjectPoliciesCmd.Flags().String("next-token", "", "The pagination token.")
		clouddirectory_listObjectPoliciesCmd.Flags().String("object-reference", "", "Reference that identifies the object for which policies will be listed.")
		clouddirectory_listObjectPoliciesCmd.MarkFlagRequired("directory-arn")
		clouddirectory_listObjectPoliciesCmd.MarkFlagRequired("object-reference")
	})
	clouddirectoryCmd.AddCommand(clouddirectory_listObjectPoliciesCmd)
}
