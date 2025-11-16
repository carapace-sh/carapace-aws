package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var macie2_listCustomDataIdentifiersCmd = &cobra.Command{
	Use:   "list-custom-data-identifiers",
	Short: "Retrieves a subset of information about the custom data identifiers for an account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(macie2_listCustomDataIdentifiersCmd).Standalone()

	macie2_listCustomDataIdentifiersCmd.Flags().String("max-results", "", "The maximum number of items to include in each page of the response.")
	macie2_listCustomDataIdentifiersCmd.Flags().String("next-token", "", "The nextToken string that specifies which page of results to return in a paginated response.")
	macie2Cmd.AddCommand(macie2_listCustomDataIdentifiersCmd)
}
