package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amp_putResourcePolicyCmd = &cobra.Command{
	Use:   "put-resource-policy",
	Short: "Creates or updates a resource-based policy for an Amazon Managed Service for Prometheus workspace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amp_putResourcePolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(amp_putResourcePolicyCmd).Standalone()

		amp_putResourcePolicyCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the request is safe to retry (idempotent).")
		amp_putResourcePolicyCmd.Flags().String("policy-document", "", "The JSON policy document to use as the resource-based policy.")
		amp_putResourcePolicyCmd.Flags().String("revision-id", "", "The revision ID of the policy to update.")
		amp_putResourcePolicyCmd.Flags().String("workspace-id", "", "The ID of the workspace to attach the resource-based policy to.")
		amp_putResourcePolicyCmd.MarkFlagRequired("policy-document")
		amp_putResourcePolicyCmd.MarkFlagRequired("workspace-id")
	})
	ampCmd.AddCommand(amp_putResourcePolicyCmd)
}
