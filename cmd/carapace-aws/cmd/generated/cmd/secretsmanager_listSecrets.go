package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var secretsmanager_listSecretsCmd = &cobra.Command{
	Use:   "list-secrets",
	Short: "Lists the secrets that are stored by Secrets Manager in the Amazon Web Services account, not including secrets that are marked for deletion.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(secretsmanager_listSecretsCmd).Standalone()

	secretsmanager_listSecretsCmd.Flags().String("filters", "", "The filters to apply to the list of secrets.")
	secretsmanager_listSecretsCmd.Flags().String("include-planned-deletion", "", "Specifies whether to include secrets scheduled for deletion.")
	secretsmanager_listSecretsCmd.Flags().String("max-results", "", "The number of results to include in the response.")
	secretsmanager_listSecretsCmd.Flags().String("next-token", "", "A token that indicates where the output should continue from, if a previous call did not show all results.")
	secretsmanager_listSecretsCmd.Flags().String("sort-order", "", "Secrets are listed by `CreatedDate`.")
	secretsmanagerCmd.AddCommand(secretsmanager_listSecretsCmd)
}
