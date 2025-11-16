package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var secretsmanager_getSecretValueCmd = &cobra.Command{
	Use:   "get-secret-value",
	Short: "Retrieves the contents of the encrypted fields `SecretString` or `SecretBinary` from the specified version of a secret, whichever contains content.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(secretsmanager_getSecretValueCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(secretsmanager_getSecretValueCmd).Standalone()

		secretsmanager_getSecretValueCmd.Flags().String("secret-id", "", "The ARN or name of the secret to retrieve.")
		secretsmanager_getSecretValueCmd.Flags().String("version-id", "", "The unique identifier of the version of the secret to retrieve.")
		secretsmanager_getSecretValueCmd.Flags().String("version-stage", "", "The staging label of the version of the secret to retrieve.")
		secretsmanager_getSecretValueCmd.MarkFlagRequired("secret-id")
	})
	secretsmanagerCmd.AddCommand(secretsmanager_getSecretValueCmd)
}
