package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workdocs_describeNotificationSubscriptionsCmd = &cobra.Command{
	Use:   "describe-notification-subscriptions",
	Short: "Lists the specified notification subscriptions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workdocs_describeNotificationSubscriptionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workdocs_describeNotificationSubscriptionsCmd).Standalone()

		workdocs_describeNotificationSubscriptionsCmd.Flags().String("limit", "", "The maximum number of items to return with this call.")
		workdocs_describeNotificationSubscriptionsCmd.Flags().String("marker", "", "The marker for the next set of results.")
		workdocs_describeNotificationSubscriptionsCmd.Flags().String("organization-id", "", "The ID of the organization.")
		workdocs_describeNotificationSubscriptionsCmd.MarkFlagRequired("organization-id")
	})
	workdocsCmd.AddCommand(workdocs_describeNotificationSubscriptionsCmd)
}
