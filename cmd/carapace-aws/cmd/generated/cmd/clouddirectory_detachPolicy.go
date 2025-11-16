package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var clouddirectory_detachPolicyCmd = &cobra.Command{
	Use:   "detach-policy",
	Short: "Detaches a policy from an object.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clouddirectory_detachPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(clouddirectory_detachPolicyCmd).Standalone()

		clouddirectory_detachPolicyCmd.Flags().String("directory-arn", "", "The Amazon Resource Name (ARN) that is associated with the [Directory]() where both objects reside.")
		clouddirectory_detachPolicyCmd.Flags().String("object-reference", "", "Reference that identifies the object whose policy object will be detached.")
		clouddirectory_detachPolicyCmd.Flags().String("policy-reference", "", "Reference that identifies the policy object.")
		clouddirectory_detachPolicyCmd.MarkFlagRequired("directory-arn")
		clouddirectory_detachPolicyCmd.MarkFlagRequired("object-reference")
		clouddirectory_detachPolicyCmd.MarkFlagRequired("policy-reference")
	})
	clouddirectoryCmd.AddCommand(clouddirectory_detachPolicyCmd)
}
