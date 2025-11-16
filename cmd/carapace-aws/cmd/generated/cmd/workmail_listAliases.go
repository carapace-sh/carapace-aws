package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_listAliasesCmd = &cobra.Command{
	Use:   "list-aliases",
	Short: "Creates a paginated call to list the aliases associated with a given entity.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_listAliasesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workmail_listAliasesCmd).Standalone()

		workmail_listAliasesCmd.Flags().String("entity-id", "", "The identifier for the entity for which to list the aliases.")
		workmail_listAliasesCmd.Flags().String("max-results", "", "The maximum number of results to return in a single call.")
		workmail_listAliasesCmd.Flags().String("next-token", "", "The token to use to retrieve the next page of results.")
		workmail_listAliasesCmd.Flags().String("organization-id", "", "The identifier for the organization under which the entity exists.")
		workmail_listAliasesCmd.MarkFlagRequired("entity-id")
		workmail_listAliasesCmd.MarkFlagRequired("organization-id")
	})
	workmailCmd.AddCommand(workmail_listAliasesCmd)
}
