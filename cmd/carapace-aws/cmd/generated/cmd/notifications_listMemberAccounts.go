package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var notifications_listMemberAccountsCmd = &cobra.Command{
	Use:   "list-member-accounts",
	Short: "Returns a list of member accounts associated with a notification configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(notifications_listMemberAccountsCmd).Standalone()

	notifications_listMemberAccountsCmd.Flags().String("max-results", "", "The maximum number of results to return in a single call.")
	notifications_listMemberAccountsCmd.Flags().String("member-account", "", "The member account identifier used to filter the results.")
	notifications_listMemberAccountsCmd.Flags().String("next-token", "", "The token for the next page of results.")
	notifications_listMemberAccountsCmd.Flags().String("notification-configuration-arn", "", "The Amazon Resource Name (ARN) of the notification configuration used to filter the member accounts.")
	notifications_listMemberAccountsCmd.Flags().String("organizational-unit-id", "", "The organizational unit ID used to filter the member accounts.")
	notifications_listMemberAccountsCmd.Flags().String("status", "", "The status used to filter the member accounts.")
	notifications_listMemberAccountsCmd.MarkFlagRequired("notification-configuration-arn")
	notificationsCmd.AddCommand(notifications_listMemberAccountsCmd)
}
