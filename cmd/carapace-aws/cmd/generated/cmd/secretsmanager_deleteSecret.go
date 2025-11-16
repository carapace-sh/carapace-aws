package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var secretsmanager_deleteSecretCmd = &cobra.Command{
	Use:   "delete-secret",
	Short: "Deletes a secret and all of its versions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(secretsmanager_deleteSecretCmd).Standalone()

	secretsmanager_deleteSecretCmd.Flags().String("force-delete-without-recovery", "", "Specifies whether to delete the secret without any recovery window.")
	secretsmanager_deleteSecretCmd.Flags().String("recovery-window-in-days", "", "The number of days from 7 to 30 that Secrets Manager waits before permanently deleting the secret.")
	secretsmanager_deleteSecretCmd.Flags().String("secret-id", "", "The ARN or name of the secret to delete.")
	secretsmanager_deleteSecretCmd.MarkFlagRequired("secret-id")
	secretsmanagerCmd.AddCommand(secretsmanager_deleteSecretCmd)
}
