package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var secretsmanager_cancelRotateSecretCmd = &cobra.Command{
	Use:   "cancel-rotate-secret",
	Short: "Turns off automatic rotation, and if a rotation is currently in progress, cancels the rotation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(secretsmanager_cancelRotateSecretCmd).Standalone()

	secretsmanager_cancelRotateSecretCmd.Flags().String("secret-id", "", "The ARN or name of the secret.")
	secretsmanager_cancelRotateSecretCmd.MarkFlagRequired("secret-id")
	secretsmanagerCmd.AddCommand(secretsmanager_cancelRotateSecretCmd)
}
