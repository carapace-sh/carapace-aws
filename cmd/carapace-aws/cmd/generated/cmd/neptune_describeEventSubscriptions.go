package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptune_describeEventSubscriptionsCmd = &cobra.Command{
	Use:   "describe-event-subscriptions",
	Short: "Lists all the subscription descriptions for a customer account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptune_describeEventSubscriptionsCmd).Standalone()

	neptune_describeEventSubscriptionsCmd.Flags().String("filters", "", "This parameter is not currently supported.")
	neptune_describeEventSubscriptionsCmd.Flags().String("marker", "", "An optional pagination token provided by a previous DescribeOrderableDBInstanceOptions request.")
	neptune_describeEventSubscriptionsCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
	neptune_describeEventSubscriptionsCmd.Flags().String("subscription-name", "", "The name of the event notification subscription you want to describe.")
	neptuneCmd.AddCommand(neptune_describeEventSubscriptionsCmd)
}
