package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var secretsmanager_validateResourcePolicyCmd = &cobra.Command{
	Use:   "validate-resource-policy",
	Short: "Validates that a resource policy does not grant a wide range of principals access to your secret.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(secretsmanager_validateResourcePolicyCmd).Standalone()

	secretsmanager_validateResourcePolicyCmd.Flags().String("resource-policy", "", "A JSON-formatted string that contains an Amazon Web Services resource-based policy.")
	secretsmanager_validateResourcePolicyCmd.Flags().String("secret-id", "", "The ARN or name of the secret with the resource-based policy you want to validate.")
	secretsmanager_validateResourcePolicyCmd.MarkFlagRequired("resource-policy")
	secretsmanagerCmd.AddCommand(secretsmanager_validateResourcePolicyCmd)
}
