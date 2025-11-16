package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var organizations_attachPolicyCmd = &cobra.Command{
	Use:   "attach-policy",
	Short: "Attaches a policy to a root, an organizational unit (OU), or an individual account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(organizations_attachPolicyCmd).Standalone()

	organizations_attachPolicyCmd.Flags().String("policy-id", "", "The unique identifier (ID) of the policy that you want to attach to the target.")
	organizations_attachPolicyCmd.Flags().String("target-id", "", "The unique identifier (ID) of the root, OU, or account that you want to attach the policy to.")
	organizations_attachPolicyCmd.MarkFlagRequired("policy-id")
	organizations_attachPolicyCmd.MarkFlagRequired("target-id")
	organizationsCmd.AddCommand(organizations_attachPolicyCmd)
}
