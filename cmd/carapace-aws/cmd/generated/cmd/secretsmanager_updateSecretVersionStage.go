package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var secretsmanager_updateSecretVersionStageCmd = &cobra.Command{
	Use:   "update-secret-version-stage",
	Short: "Modifies the staging labels attached to a version of a secret.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(secretsmanager_updateSecretVersionStageCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(secretsmanager_updateSecretVersionStageCmd).Standalone()

		secretsmanager_updateSecretVersionStageCmd.Flags().String("move-to-version-id", "", "The ID of the version to add the staging label to.")
		secretsmanager_updateSecretVersionStageCmd.Flags().String("remove-from-version-id", "", "The ID of the version that the staging label is to be removed from.")
		secretsmanager_updateSecretVersionStageCmd.Flags().String("secret-id", "", "The ARN or the name of the secret with the version and staging labelsto modify.")
		secretsmanager_updateSecretVersionStageCmd.Flags().String("version-stage", "", "The staging label to add to this version.")
		secretsmanager_updateSecretVersionStageCmd.MarkFlagRequired("secret-id")
		secretsmanager_updateSecretVersionStageCmd.MarkFlagRequired("version-stage")
	})
	secretsmanagerCmd.AddCommand(secretsmanager_updateSecretVersionStageCmd)
}
