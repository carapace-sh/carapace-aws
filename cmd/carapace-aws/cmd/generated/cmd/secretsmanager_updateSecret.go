package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var secretsmanager_updateSecretCmd = &cobra.Command{
	Use:   "update-secret",
	Short: "Modifies the details of a secret, including metadata and the secret value.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(secretsmanager_updateSecretCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(secretsmanager_updateSecretCmd).Standalone()

		secretsmanager_updateSecretCmd.Flags().String("client-request-token", "", "If you include `SecretString` or `SecretBinary`, then Secrets Manager creates a new version for the secret, and this parameter specifies the unique identifier for the new version.")
		secretsmanager_updateSecretCmd.Flags().String("description", "", "The description of the secret.")
		secretsmanager_updateSecretCmd.Flags().String("kms-key-id", "", "The ARN, key ID, or alias of the KMS key that Secrets Manager uses to encrypt new secret versions as well as any existing versions with the staging labels `AWSCURRENT`, `AWSPENDING`, or `AWSPREVIOUS`.")
		secretsmanager_updateSecretCmd.Flags().String("secret-binary", "", "The binary data to encrypt and store in the new version of the secret.")
		secretsmanager_updateSecretCmd.Flags().String("secret-id", "", "The ARN or name of the secret.")
		secretsmanager_updateSecretCmd.Flags().String("secret-string", "", "The text data to encrypt and store in the new version of the secret.")
		secretsmanager_updateSecretCmd.MarkFlagRequired("secret-id")
	})
	secretsmanagerCmd.AddCommand(secretsmanager_updateSecretCmd)
}
