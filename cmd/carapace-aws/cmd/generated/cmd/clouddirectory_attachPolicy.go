package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var clouddirectory_attachPolicyCmd = &cobra.Command{
	Use:   "attach-policy",
	Short: "Attaches a policy object to a regular object.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clouddirectory_attachPolicyCmd).Standalone()

	clouddirectory_attachPolicyCmd.Flags().String("directory-arn", "", "The Amazon Resource Name (ARN) that is associated with the [Directory]() where both objects reside.")
	clouddirectory_attachPolicyCmd.Flags().String("object-reference", "", "The reference that identifies the object to which the policy will be attached.")
	clouddirectory_attachPolicyCmd.Flags().String("policy-reference", "", "The reference that is associated with the policy object.")
	clouddirectory_attachPolicyCmd.MarkFlagRequired("directory-arn")
	clouddirectory_attachPolicyCmd.MarkFlagRequired("object-reference")
	clouddirectory_attachPolicyCmd.MarkFlagRequired("policy-reference")
	clouddirectoryCmd.AddCommand(clouddirectory_attachPolicyCmd)
}
