package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var macie2_listClassificationScopesCmd = &cobra.Command{
	Use:   "list-classification-scopes",
	Short: "Retrieves a subset of information about the classification scope for an account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(macie2_listClassificationScopesCmd).Standalone()

	macie2_listClassificationScopesCmd.Flags().String("name", "", "The name of the classification scope to retrieve the unique identifier for.")
	macie2_listClassificationScopesCmd.Flags().String("next-token", "", "The nextToken string that specifies which page of results to return in a paginated response.")
	macie2Cmd.AddCommand(macie2_listClassificationScopesCmd)
}
