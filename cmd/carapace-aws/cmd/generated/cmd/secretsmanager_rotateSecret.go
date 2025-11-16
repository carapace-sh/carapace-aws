package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var secretsmanager_rotateSecretCmd = &cobra.Command{
	Use:   "rotate-secret",
	Short: "Configures and starts the asynchronous process of rotating the secret.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(secretsmanager_rotateSecretCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(secretsmanager_rotateSecretCmd).Standalone()

		secretsmanager_rotateSecretCmd.Flags().String("client-request-token", "", "A unique identifier for the new version of the secret.")
		secretsmanager_rotateSecretCmd.Flags().String("rotate-immediately", "", "Specifies whether to rotate the secret immediately or wait until the next scheduled rotation window.")
		secretsmanager_rotateSecretCmd.Flags().String("rotation-lambda-arn", "", "For secrets that use a Lambda rotation function to rotate, the ARN of the Lambda rotation function.")
		secretsmanager_rotateSecretCmd.Flags().String("rotation-rules", "", "A structure that defines the rotation configuration for this secret.")
		secretsmanager_rotateSecretCmd.Flags().String("secret-id", "", "The ARN or name of the secret to rotate.")
		secretsmanager_rotateSecretCmd.MarkFlagRequired("secret-id")
	})
	secretsmanagerCmd.AddCommand(secretsmanager_rotateSecretCmd)
}
