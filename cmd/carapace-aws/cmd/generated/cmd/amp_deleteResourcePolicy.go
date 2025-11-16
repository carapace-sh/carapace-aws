package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amp_deleteResourcePolicyCmd = &cobra.Command{
	Use:   "delete-resource-policy",
	Short: "Deletes the resource-based policy attached to an Amazon Managed Service for Prometheus workspace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amp_deleteResourcePolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(amp_deleteResourcePolicyCmd).Standalone()

		amp_deleteResourcePolicyCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the request is safe to retry (idempotent).")
		amp_deleteResourcePolicyCmd.Flags().String("revision-id", "", "The revision ID of the policy to delete.")
		amp_deleteResourcePolicyCmd.Flags().String("workspace-id", "", "The ID of the workspace from which to delete the resource-based policy.")
		amp_deleteResourcePolicyCmd.MarkFlagRequired("workspace-id")
	})
	ampCmd.AddCommand(amp_deleteResourcePolicyCmd)
}
