package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var secretsmanager_createSecretCmd = &cobra.Command{
	Use:   "create-secret",
	Short: "Creates a new secret.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(secretsmanager_createSecretCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(secretsmanager_createSecretCmd).Standalone()

		secretsmanager_createSecretCmd.Flags().String("add-replica-regions", "", "A list of Regions and KMS keys to replicate secrets.")
		secretsmanager_createSecretCmd.Flags().String("client-request-token", "", "If you include `SecretString` or `SecretBinary`, then Secrets Manager creates an initial version for the secret, and this parameter specifies the unique identifier for the new version.")
		secretsmanager_createSecretCmd.Flags().String("description", "", "The description of the secret.")
		secretsmanager_createSecretCmd.Flags().String("force-overwrite-replica-secret", "", "Specifies whether to overwrite a secret with the same name in the destination Region.")
		secretsmanager_createSecretCmd.Flags().String("kms-key-id", "", "The ARN, key ID, or alias of the KMS key that Secrets Manager uses to encrypt the secret value in the secret.")
		secretsmanager_createSecretCmd.Flags().String("name", "", "The name of the new secret.")
		secretsmanager_createSecretCmd.Flags().String("secret-binary", "", "The binary data to encrypt and store in the new version of the secret.")
		secretsmanager_createSecretCmd.Flags().String("secret-string", "", "The text data to encrypt and store in this new version of the secret.")
		secretsmanager_createSecretCmd.Flags().String("tags", "", "A list of tags to attach to the secret.")
		secretsmanager_createSecretCmd.MarkFlagRequired("name")
	})
	secretsmanagerCmd.AddCommand(secretsmanager_createSecretCmd)
}
