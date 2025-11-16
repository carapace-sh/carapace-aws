package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var secretsmanager_restoreSecretCmd = &cobra.Command{
	Use:   "restore-secret",
	Short: "Cancels the scheduled deletion of a secret by removing the `DeletedDate` time stamp.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(secretsmanager_restoreSecretCmd).Standalone()

	secretsmanager_restoreSecretCmd.Flags().String("secret-id", "", "The ARN or name of the secret to restore.")
	secretsmanager_restoreSecretCmd.MarkFlagRequired("secret-id")
	secretsmanagerCmd.AddCommand(secretsmanager_restoreSecretCmd)
}
