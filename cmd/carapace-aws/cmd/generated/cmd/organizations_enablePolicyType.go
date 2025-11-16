package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var organizations_enablePolicyTypeCmd = &cobra.Command{
	Use:   "enable-policy-type",
	Short: "Enables a policy type in a root.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(organizations_enablePolicyTypeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(organizations_enablePolicyTypeCmd).Standalone()

		organizations_enablePolicyTypeCmd.Flags().String("policy-type", "", "The policy type that you want to enable.")
		organizations_enablePolicyTypeCmd.Flags().String("root-id", "", "The unique identifier (ID) of the root in which you want to enable a policy type.")
		organizations_enablePolicyTypeCmd.MarkFlagRequired("policy-type")
		organizations_enablePolicyTypeCmd.MarkFlagRequired("root-id")
	})
	organizationsCmd.AddCommand(organizations_enablePolicyTypeCmd)
}
