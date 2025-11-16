package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var secretsmanager_putSecretValueCmd = &cobra.Command{
	Use:   "put-secret-value",
	Short: "Creates a new version with a new encrypted secret value and attaches it to the secret.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(secretsmanager_putSecretValueCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(secretsmanager_putSecretValueCmd).Standalone()

		secretsmanager_putSecretValueCmd.Flags().String("client-request-token", "", "A unique identifier for the new version of the secret.")
		secretsmanager_putSecretValueCmd.Flags().String("rotation-token", "", "A unique identifier that indicates the source of the request.")
		secretsmanager_putSecretValueCmd.Flags().String("secret-binary", "", "The binary data to encrypt and store in the new version of the secret.")
		secretsmanager_putSecretValueCmd.Flags().String("secret-id", "", "The ARN or name of the secret to add a new version to.")
		secretsmanager_putSecretValueCmd.Flags().String("secret-string", "", "The text to encrypt and store in the new version of the secret.")
		secretsmanager_putSecretValueCmd.Flags().String("version-stages", "", "A list of staging labels to attach to this version of the secret.")
		secretsmanager_putSecretValueCmd.MarkFlagRequired("secret-id")
	})
	secretsmanagerCmd.AddCommand(secretsmanager_putSecretValueCmd)
}
