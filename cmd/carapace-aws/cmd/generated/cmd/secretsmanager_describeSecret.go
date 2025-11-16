package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var secretsmanager_describeSecretCmd = &cobra.Command{
	Use:   "describe-secret",
	Short: "Retrieves the details of a secret.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(secretsmanager_describeSecretCmd).Standalone()

	secretsmanager_describeSecretCmd.Flags().String("secret-id", "", "The ARN or name of the secret.")
	secretsmanager_describeSecretCmd.MarkFlagRequired("secret-id")
	secretsmanagerCmd.AddCommand(secretsmanager_describeSecretCmd)
}
