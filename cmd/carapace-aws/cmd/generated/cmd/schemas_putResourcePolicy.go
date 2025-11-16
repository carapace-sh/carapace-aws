package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var schemas_putResourcePolicyCmd = &cobra.Command{
	Use:   "put-resource-policy",
	Short: "The name of the policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(schemas_putResourcePolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(schemas_putResourcePolicyCmd).Standalone()

		schemas_putResourcePolicyCmd.Flags().String("policy", "", "The resource-based policy.")
		schemas_putResourcePolicyCmd.Flags().String("registry-name", "", "The name of the registry.")
		schemas_putResourcePolicyCmd.Flags().String("revision-id", "", "The revision ID of the policy.")
		schemas_putResourcePolicyCmd.MarkFlagRequired("policy")
	})
	schemasCmd.AddCommand(schemas_putResourcePolicyCmd)
}
