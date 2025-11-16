package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var docdb_describeEventSubscriptionsCmd = &cobra.Command{
	Use:   "describe-event-subscriptions",
	Short: "Lists all the subscription descriptions for a customer account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docdb_describeEventSubscriptionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(docdb_describeEventSubscriptionsCmd).Standalone()

		docdb_describeEventSubscriptionsCmd.Flags().String("filters", "", "This parameter is not currently supported.")
		docdb_describeEventSubscriptionsCmd.Flags().String("marker", "", "An optional pagination token provided by a previous request.")
		docdb_describeEventSubscriptionsCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
		docdb_describeEventSubscriptionsCmd.Flags().String("subscription-name", "", "The name of the Amazon DocumentDB event notification subscription that you want to describe.")
	})
	docdbCmd.AddCommand(docdb_describeEventSubscriptionsCmd)
}
