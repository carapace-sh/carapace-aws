package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var macie2_listMembersCmd = &cobra.Command{
	Use:   "list-members",
	Short: "Retrieves information about the accounts that are associated with an Amazon Macie administrator account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(macie2_listMembersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(macie2_listMembersCmd).Standalone()

		macie2_listMembersCmd.Flags().String("max-results", "", "The maximum number of items to include in each page of a paginated response.")
		macie2_listMembersCmd.Flags().String("next-token", "", "The nextToken string that specifies which page of results to return in a paginated response.")
		macie2_listMembersCmd.Flags().String("only-associated", "", "Specifies which accounts to include in the response, based on the status of an account's relationship with the administrator account.")
	})
	macie2Cmd.AddCommand(macie2_listMembersCmd)
}
