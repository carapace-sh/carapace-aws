package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var secretsmanager_putResourcePolicyCmd = &cobra.Command{
	Use:   "put-resource-policy",
	Short: "Attaches a resource-based permission policy to a secret.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(secretsmanager_putResourcePolicyCmd).Standalone()

	secretsmanager_putResourcePolicyCmd.Flags().String("block-public-policy", "", "Specifies whether to block resource-based policies that allow broad access to the secret, for example those that use a wildcard for the principal.")
	secretsmanager_putResourcePolicyCmd.Flags().String("resource-policy", "", "A JSON-formatted string for an Amazon Web Services resource-based policy.")
	secretsmanager_putResourcePolicyCmd.Flags().String("secret-id", "", "The ARN or name of the secret to attach the resource-based policy.")
	secretsmanager_putResourcePolicyCmd.MarkFlagRequired("resource-policy")
	secretsmanager_putResourcePolicyCmd.MarkFlagRequired("secret-id")
	secretsmanagerCmd.AddCommand(secretsmanager_putResourcePolicyCmd)
}
