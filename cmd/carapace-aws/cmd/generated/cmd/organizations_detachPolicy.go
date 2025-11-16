package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var organizations_detachPolicyCmd = &cobra.Command{
	Use:   "detach-policy",
	Short: "Detaches a policy from a target root, organizational unit (OU), or account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(organizations_detachPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(organizations_detachPolicyCmd).Standalone()

		organizations_detachPolicyCmd.Flags().String("policy-id", "", "The unique identifier (ID) of the policy you want to detach.")
		organizations_detachPolicyCmd.Flags().String("target-id", "", "The unique identifier (ID) of the root, OU, or account that you want to detach the policy from.")
		organizations_detachPolicyCmd.MarkFlagRequired("policy-id")
		organizations_detachPolicyCmd.MarkFlagRequired("target-id")
	})
	organizationsCmd.AddCommand(organizations_detachPolicyCmd)
}
