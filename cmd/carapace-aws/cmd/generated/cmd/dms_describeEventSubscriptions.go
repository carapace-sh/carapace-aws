package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_describeEventSubscriptionsCmd = &cobra.Command{
	Use:   "describe-event-subscriptions",
	Short: "Lists all the event subscriptions for a customer account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_describeEventSubscriptionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dms_describeEventSubscriptionsCmd).Standalone()

		dms_describeEventSubscriptionsCmd.Flags().String("filters", "", "Filters applied to event subscriptions.")
		dms_describeEventSubscriptionsCmd.Flags().String("marker", "", "An optional pagination token provided by a previous request.")
		dms_describeEventSubscriptionsCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
		dms_describeEventSubscriptionsCmd.Flags().String("subscription-name", "", "The name of the DMS event subscription to be described.")
	})
	dmsCmd.AddCommand(dms_describeEventSubscriptionsCmd)
}
