package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_listLogGroupsCmd = &cobra.Command{
	Use:   "list-log-groups",
	Short: "Returns a list of log groups in the Region in your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_listLogGroupsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(logs_listLogGroupsCmd).Standalone()

		logs_listLogGroupsCmd.Flags().String("account-identifiers", "", "When `includeLinkedAccounts` is set to `true`, use this parameter to specify the list of accounts to search.")
		logs_listLogGroupsCmd.Flags().String("include-linked-accounts", "", "If you are using a monitoring account, set this to `true` to have the operation return log groups in the accounts listed in `accountIdentifiers`.")
		logs_listLogGroupsCmd.Flags().String("limit", "", "The maximum number of log groups to return.")
		logs_listLogGroupsCmd.Flags().String("log-group-class", "", "Use this parameter to limit the results to only those log groups in the specified log group class.")
		logs_listLogGroupsCmd.Flags().String("log-group-name-pattern", "", "Use this parameter to limit the returned log groups to only those with names that match the pattern that you specify.")
		logs_listLogGroupsCmd.Flags().String("next-token", "", "")
	})
	logsCmd.AddCommand(logs_listLogGroupsCmd)
}
