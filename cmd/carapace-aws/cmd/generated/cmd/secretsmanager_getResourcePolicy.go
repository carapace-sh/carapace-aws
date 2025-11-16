package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var secretsmanager_getResourcePolicyCmd = &cobra.Command{
	Use:   "get-resource-policy",
	Short: "Retrieves the JSON text of the resource-based policy document attached to the secret.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(secretsmanager_getResourcePolicyCmd).Standalone()

	secretsmanager_getResourcePolicyCmd.Flags().String("secret-id", "", "The ARN or name of the secret to retrieve the attached resource-based policy for.")
	secretsmanager_getResourcePolicyCmd.MarkFlagRequired("secret-id")
	secretsmanagerCmd.AddCommand(secretsmanager_getResourcePolicyCmd)
}
