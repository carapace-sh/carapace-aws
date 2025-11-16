package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var secretsmanager_batchGetSecretValueCmd = &cobra.Command{
	Use:   "batch-get-secret-value",
	Short: "Retrieves the contents of the encrypted fields `SecretString` or `SecretBinary` for up to 20 secrets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(secretsmanager_batchGetSecretValueCmd).Standalone()

	secretsmanager_batchGetSecretValueCmd.Flags().String("filters", "", "The filters to choose which secrets to retrieve.")
	secretsmanager_batchGetSecretValueCmd.Flags().String("max-results", "", "The number of results to include in the response.")
	secretsmanager_batchGetSecretValueCmd.Flags().String("next-token", "", "A token that indicates where the output should continue from, if a previous call did not show all results.")
	secretsmanager_batchGetSecretValueCmd.Flags().String("secret-id-list", "", "The ARN or names of the secrets to retrieve.")
	secretsmanagerCmd.AddCommand(secretsmanager_batchGetSecretValueCmd)
}
