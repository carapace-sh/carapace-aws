package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_describeLogGroupsCmd = &cobra.Command{
	Use:   "describe-log-groups",
	Short: "Returns information about log groups.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_describeLogGroupsCmd).Standalone()

	logs_describeLogGroupsCmd.Flags().String("account-identifiers", "", "When `includeLinkedAccounts` is set to `true`, use this parameter to specify the list of accounts to search.")
	logs_describeLogGroupsCmd.Flags().String("include-linked-accounts", "", "If you are using a monitoring account, set this to `true` to have the operation return log groups in the accounts listed in `accountIdentifiers`.")
	logs_describeLogGroupsCmd.Flags().String("limit", "", "The maximum number of items returned.")
	logs_describeLogGroupsCmd.Flags().String("log-group-class", "", "Use this parameter to limit the results to only those log groups in the specified log group class.")
	logs_describeLogGroupsCmd.Flags().String("log-group-identifiers", "", "Use this array to filter the list of log groups returned.")
	logs_describeLogGroupsCmd.Flags().String("log-group-name-pattern", "", "If you specify a string for this parameter, the operation returns only log groups that have names that match the string based on a case-sensitive substring search.")
	logs_describeLogGroupsCmd.Flags().String("log-group-name-prefix", "", "The prefix to match.")
	logs_describeLogGroupsCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
	logsCmd.AddCommand(logs_describeLogGroupsCmd)
}
