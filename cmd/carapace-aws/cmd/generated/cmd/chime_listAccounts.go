package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chime_listAccountsCmd = &cobra.Command{
	Use:   "list-accounts",
	Short: "Lists the Amazon Chime accounts under the administrator's AWS account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chime_listAccountsCmd).Standalone()

	chime_listAccountsCmd.Flags().String("max-results", "", "The maximum number of results to return in a single call.")
	chime_listAccountsCmd.Flags().String("name", "", "Amazon Chime account name prefix with which to filter results.")
	chime_listAccountsCmd.Flags().String("next-token", "", "The token to use to retrieve the next page of results.")
	chime_listAccountsCmd.Flags().String("user-email", "", "User email address with which to filter results.")
	chimeCmd.AddCommand(chime_listAccountsCmd)
}
