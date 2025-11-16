package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var secretsmanager_listSecretVersionIdsCmd = &cobra.Command{
	Use:   "list-secret-version-ids",
	Short: "Lists the versions of a secret.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(secretsmanager_listSecretVersionIdsCmd).Standalone()

	secretsmanager_listSecretVersionIdsCmd.Flags().String("include-deprecated", "", "Specifies whether to include versions of secrets that don't have any staging labels attached to them.")
	secretsmanager_listSecretVersionIdsCmd.Flags().String("max-results", "", "The number of results to include in the response.")
	secretsmanager_listSecretVersionIdsCmd.Flags().String("next-token", "", "A token that indicates where the output should continue from, if a previous call did not show all results.")
	secretsmanager_listSecretVersionIdsCmd.Flags().String("secret-id", "", "The ARN or name of the secret whose versions you want to list.")
	secretsmanager_listSecretVersionIdsCmd.MarkFlagRequired("secret-id")
	secretsmanagerCmd.AddCommand(secretsmanager_listSecretVersionIdsCmd)
}
