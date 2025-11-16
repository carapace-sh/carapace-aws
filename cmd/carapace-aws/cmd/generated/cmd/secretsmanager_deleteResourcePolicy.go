package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var secretsmanager_deleteResourcePolicyCmd = &cobra.Command{
	Use:   "delete-resource-policy",
	Short: "Deletes the resource-based permission policy attached to the secret.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(secretsmanager_deleteResourcePolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(secretsmanager_deleteResourcePolicyCmd).Standalone()

		secretsmanager_deleteResourcePolicyCmd.Flags().String("secret-id", "", "The ARN or name of the secret to delete the attached resource-based policy for.")
		secretsmanager_deleteResourcePolicyCmd.MarkFlagRequired("secret-id")
	})
	secretsmanagerCmd.AddCommand(secretsmanager_deleteResourcePolicyCmd)
}
