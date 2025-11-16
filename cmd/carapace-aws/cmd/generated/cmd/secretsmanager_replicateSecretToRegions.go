package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var secretsmanager_replicateSecretToRegionsCmd = &cobra.Command{
	Use:   "replicate-secret-to-regions",
	Short: "Replicates the secret to a new Regions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(secretsmanager_replicateSecretToRegionsCmd).Standalone()

	secretsmanager_replicateSecretToRegionsCmd.Flags().String("add-replica-regions", "", "A list of Regions in which to replicate the secret.")
	secretsmanager_replicateSecretToRegionsCmd.Flags().String("force-overwrite-replica-secret", "", "Specifies whether to overwrite a secret with the same name in the destination Region.")
	secretsmanager_replicateSecretToRegionsCmd.Flags().String("secret-id", "", "The ARN or name of the secret to replicate.")
	secretsmanager_replicateSecretToRegionsCmd.MarkFlagRequired("add-replica-regions")
	secretsmanager_replicateSecretToRegionsCmd.MarkFlagRequired("secret-id")
	secretsmanagerCmd.AddCommand(secretsmanager_replicateSecretToRegionsCmd)
}
